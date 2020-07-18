from aws_cdk import (
    # aws_s3 as s3,
    aws_cognito as cognito,
    core
)
import logging

COGNITO_CUSTOM_DOMAIN="service-sandbox00"

class DataStack(core.Stack):

    def __init__(self, scope: core.Construct, id: str, **kwargs) -> None:
        super().__init__(scope, id, **kwargs)

        userpool = cognito.UserPool(self, 'service-userpool', 
            user_pool_name='service-userpool',
        )
        
        userpool.add_client("customer-app-client")
        user_pool_custom_domain = cognito.CfnUserPoolDomain(
            self,
            "CustomDomain",
            domain=COGNITO_CUSTOM_DOMAIN,
            user_pool_id=userpool.user_pool_id
        )

        cognito.CfnIdentityPool(self, 'service-identitypool', 
            allow_unauthenticated_identities=True
        )

        # Identity Pool higher-order not implemented yet... 
        # easier to manually configure IdentityPool roles on Console
        #auth_role_attachment = cognito.CfnIdentityPoolRoleAttachment(self, 
        #    id="service-cognito-auth", identity_pool_id=identitypool.id)
        #unauth_role_attachment = cognito.CfnIdentityPoolRoleAttachment(self, 
        #   id="service-cognito-unnauth", identity_pool_id=identitypool.id)

