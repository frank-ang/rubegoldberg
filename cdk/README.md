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
