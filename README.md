# Demo service resources

The Swish Machine: 70 Step Basketball Trickshot (Rube Goldberg Machine)
https://www.youtube.com/watch?v=Ss-P4qLLUyk

[![Basketball Rube Goldberg machine](http://img.youtube.com/vi/Ss-P4qLLUyk/0.jpg)](http://www.youtube.com/watch?v=Ss-P4qLLUyk "Basketball Rube Goldberg machine")

## Setup Environment Resources 

These resources can currently be created from the AWS Console:

* VPC, bastion host.
* Amazon Aurora MySQL, 
* Elasticache Redis, 
* Elasticsearch

### Setup resources in CDK.

Provisions demo resources in the **service** demo VPC. 

Currently: 
* Cognito, TODO, wire this up to something?

> Note instructions here are for MacOS. Tailor accordingly.

Install/upgrade CDK
```
cdk --version
```

1. Init CDK project
    ```
    cd data
    cdk init app --language python
    source .env/bin/activate
    pip install -r requirements.txt
    cdk ls
    ```

2. Deploy.
    ```
    cdk synth
    cdk deploy
    ```

3. Make changes.
    Install modules, e.g.
    ```
    pip install aws-cdk.aws-s3
    pip install aws-cdk.aws-elasticsearch

    pip install aws-cdk.aws-cognito
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

## Load sample data.

See readmes in the subfolders.
