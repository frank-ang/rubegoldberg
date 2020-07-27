# Fortune Service

Golang App.

## Deploy Code

1. Setup CodeDeploy
    ```
    aws deploy create-application \
    --application-name FortuneApp

    aws deploy create-deployment-group \
    --application-name FortuneApp \
    --deployment-group-name FortuneAppDeploymentGroup \
    --deployment-config-name CodeDeployDefault.OneAtATime \
    --ec2-tag-filters Key=Name,Value=fortune-app/ASG,Type=KEY_AND_VALUE \
    --service-role-arn arn:aws:iam::450428438179:role/CodeDeployServiceRole

    aws deploy push \
    --application-name FortuneApp \
    --s3-location s3://sandbox00-playground/codedeploy/FortuneApp.zip \
    --ignore-hidden-files
    
    aws deploy create-deployment \
    --application-name FortuneApp \
    --s3-location bucket=sandbox00-playground,key=codedeploy/FortuneApp.zip,bundleType=zip \
    --deployment-group-name FortuneAppDeploymentGroup \
    --deployment-config-name CodeDeployDefault.OneAtATime \
    --description "First Deploy"
    ```

## CI/CD with CodePipeline
... Setup CodePipeline
