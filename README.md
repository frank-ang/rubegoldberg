# Demo service resources

Provisions demo resources in the **service** demo VPC. 

> Note instructions here are for MacOS. Tailor accordingly.

Install/upgrade CDK
```
cdk --version
```

## Setup Data resources.

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