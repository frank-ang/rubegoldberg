#!/usr/bin/env python3
import os
from aws_cdk import core
from fortune.ec2_stack import Ec2Stack

# TODO: lookup from config file.
# Edit your configuration
VPC_ID="vpc-0aab98e7906fc81e5"
KEY_NAME="macbook2018"
INSTANCE_ROLE_ARN="arn:aws:iam::450428438179:role/myEC2Role"

env = core.Environment(
    account=os.environ["CDK_DEFAULT_ACCOUNT"],
    region=os.environ["CDK_DEFAULT_REGION"])

app = core.App()
Ec2Stack(app, id="fortune-app", vpc_id=VPC_ID, key_name=KEY_NAME, 
    instance_role_arn=INSTANCE_ROLE_ARN, env=env)
app.synth()
