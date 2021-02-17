# Setup App Resources.

> Note instructions here are for MacOS. Tailor accordingly.

## CDK setup
Install/upgrade CDK
```bash
cdk --version
```

1. Init CDK project, creates this directory
    ```bash
    mkdir cdk
    cd cdk
    cdk init app --language python
    source .env/bin/activate
    pip install -r requirements.txt
    cdk ls
    ```

2. Deploy Load Balancer and Autoscaling Group
    ```bash
    # set CDK_DEFAULT_ACCOUNT and CDK_DEFAULT_REGION, create a small script for convenience.
    ./setenv.gitignore
    # deploy
    cdk synth
    cdk deploy
    ```

3. Make changes.
    Install modules, e.g.
    
    ```bash
    pip install aws-cdk.aws-s3
    pip install aws-cdk.aws-cognito
    pip install aws-cdk.aws-ec2
    pip install aws-cdk.aws_autoscaling
    pip install aws-cdk.aws_elasticloadbalancingv2
    pip install aws-cdk.aws-iam

    ```
    Make code changes to the stacks

4. Deploy changes.
    ```
    cdk diff
    cdk synth
    cdk deploy
    ```

5. Destroy.

    ```
    cdk destroy
    ```

## CloudWatch Agent

### Configure the CW Agent Config file on 1 host. Use the wizard.

```bash
sudo /opt/aws/amazon-cloudwatch-agent/bin/amazon-cloudwatch-agent-config-wizard

# configure Advanced metrics.
# watch log file: /opt/fortune/fortune.out
# Writes to /opt/aws/amazon-cloudwatch-agent/bin/config.json
# Writes to SSM Parameter Store, temporarily attach IAM policy CloudWatchAgentAdminPolicy
# Test Start the CloudWatch Agent Using the Command Line
sudo /opt/aws/amazon-cloudwatch-agent/bin/amazon-cloudwatch-agent-ctl -a fetch-config -m ec2 -s -c ssm:AmazonCloudWatch-linux-fortune
```
