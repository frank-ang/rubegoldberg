#!/bin/bash

source ../config/config.gitignore

username="$COGNITO_USERNAME"
password="$COGNITO_PASSWORD"
clientid="$COGNITO_APP_CLIENT_ID"

token_request=`aws cognito-idp initiate-auth --auth-flow USER_PASSWORD_AUTH --output json --client-id $clientid --auth-parameters USERNAME=$username,PASSWORD=$password`
response_json=$token_request

echo "Auth Response: $response_json"

AccessToken=`echo $response_json | jq -r ".AuthenticationResult.AccessToken"`
IdToken=`echo $response_json | jq -r ".AuthenticationResult.IdToken"`
RefreshToken=`echo $response_json | jq -r ".AuthenticationResult.RefreshToken"`
ExpiresIn=`echo $response_json | jq -r ".AuthenticationResult.ExpiresIn"`

token=$AccessToken
echo $token