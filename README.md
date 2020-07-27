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

Load sample data. See readmes in [./fortune-data](./fortune-data)

### Setup resources in CDK.

Provisions demo resources in the **service** demo VPC. 

See [./cdk](./cdk)
