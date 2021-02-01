# API Gateway setup.

## Configure public API GW.

Manually setup:
### HTTP API 
* "FortuneHTTP"
* custom domain name
* VPC Link for private integration to internal NLB of fortune service.
* Create Authorization:
    Name: FortuneAuthorizer
    Identity source: $request.header.Authorization
    Issuer URL: https://cognito-idp.ap-southeast-1.amazonaws.com/ap-southeast-1_MNFgAO3r8
    app client: 71r5v6rb0te77e8lbsq2g6l6t9

### REST API 
* "fortune"
* custom domain name
* VPC Link for private integration to internal NLB of fortune service.

