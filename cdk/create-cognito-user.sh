#!/bin/bash
USERNAME=$1
PASSWORD="Foobar2020!"
USERPOOL_ID=ap-southeast-1_Ybu28PisS
APPCLIENT_ID=17g19621ndq1ifm6uur3fcirj0

if [ -z "$USERNAME" ]; then
    echo "Usage $0 [USERNAME]" 1>&2
    exit 1
fi

aws cognito-idp admin-create-user \
  --user-pool-id $USERPOOL_ID \
  --username $USERNAME
sleep 2
aws cognito-idp admin-set-user-password \
    --user-pool-id  $USERPOOL_ID \
    --username $USERNAME \
    --password $PASSWORD \
    --permanent

###  For sign-up enabled only: ###
#aws cognito-idp sign-up \
#  --client-id $APPCLIENT_ID \
#  --username $USERNAME \
#  --password $PASSWORD
#
#aws cognito-idp admin-confirm-sign-up \
#  --user-pool-id $USERPOOL_ID \
#  --username $USERNAME
