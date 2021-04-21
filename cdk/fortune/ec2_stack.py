from aws_cdk import (
    # aws_cognito as cognito,
    aws_autoscaling as autoscaling,
    aws_ec2 as ec2,
    aws_elasticloadbalancingv2 as elbv2,
    aws_iam as iam,
    aws_ssm as ssm,
    core
)
import logging

# TODO Lookup from config file.
BASTION_SECURITY_GROUP_ID="sg-0f07c71aae9d2bf97"

class Ec2Stack(core.Stack):
    '''
    EC2 Autoscaling Group and NLB
    '''
    def __init__(self, scope: core.Construct, id: str, 
    vpc_id: str, key_name: str, instance_role_arn: str, **kwargs) -> None:
        super().__init__(scope, id, **kwargs)
        vpc = ec2.Vpc.from_lookup(self, "VPC", is_default=False, vpc_id=vpc_id)
        role = iam.Role.from_role_arn(self, 'Role', instance_role_arn, mutable=False)

        user_data_script = open("./fortune/userdata.sh", "rb").read()
        user_data=ec2.UserData.for_linux()
        user_data.add_commands(str(user_data_script,'utf-8'))

        app_sg = ec2.SecurityGroup(self, id="FortuneAppSG", 
            vpc=vpc, security_group_name="FortuneAppSG",)
        bastion_sg = ec2.SecurityGroup.from_security_group_id(self, "BastionSG", security_group_id=BASTION_SECURITY_GROUP_ID)
        app_sg.add_ingress_rule(peer=bastion_sg, connection=ec2.Port.all_traffic(), description="AllowBastionSG")
        app_sg.add_ingress_rule(ec2.Peer.ipv4('10.22.0.0/16'), ec2.Port.tcp(80), description="TrustedNetwork")

        asg = autoscaling.AutoScalingGroup(
            scope=self,
            id="ASG",
            vpc=vpc,
            instance_type=ec2.InstanceType.of (
                ec2.InstanceClass.BURSTABLE3, ec2.InstanceSize.MICRO
            ),
            machine_image=ec2.AmazonLinuxImage(generation=ec2.AmazonLinuxGeneration.AMAZON_LINUX_2),
            user_data=user_data,
            role=role,
            key_name=key_name,
            min_capacity=1,
            desired_capacity=1,
            max_capacity=3,
            spot_price="0.0132", # TODO use Launch Template to setup Spot.
            security_group=app_sg,
            group_metrics=[autoscaling.GroupMetrics.all()],
            health_check=autoscaling.HealthCheck.elb(grace=core.Duration.seconds(10)),
        )

        core.Tags.of(asg).add("Name", "fortune-app/ASG", apply_to_launched_instances=True)
        core.Tags.of(asg).add("project", "fortune", apply_to_launched_instances=True)

        # private NLB
        nlb = elbv2.NetworkLoadBalancer(self, "FortunePrivateNLB", 
            cross_zone_enabled=True, 
            vpc=vpc,
            internet_facing=False)

        health_check = elbv2.HealthCheck(
            path="/",
            protocol=elbv2.ApplicationProtocol.HTTP,
            healthy_threshold_count=3,
            unhealthy_threshold_count=3,
            interval=core.Duration.seconds(10),
        )

        listener = nlb.add_listener("PrivateListener80", port=80)
        listener.add_targets("Target", port=80, targets=[asg], 
            deregistration_delay=core.Duration.seconds(10), 
            health_check=health_check, 
        )

        # asg.scale_on_outgoing_bytes(id="ScaleOnNetworkOut", target_bytes_per_second=100000)
        core.CfnOutput(self,"NetworkLoadBalancer",export_name="NetworkLoadBalancer",value=nlb.load_balancer_dns_name)

