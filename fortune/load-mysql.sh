#!/bin/bash
#
source ../config/config.gitignore

config-db-secret() {
	echo config-db-secret!
	SECRET_JSON=`aws secretsmanager get-secret-value --secret-id ${DB_SECRET_NAME} --query 'SecretString' --output text`
	DB_USERNAME=`echo $SECRET_JSON | jq -r '.username'`
	DB_PASSWORD=`echo $SECRET_JSON | jq -r '.password'`
	DB_HOST=`echo $SECRET_JSON | jq -r '.host'`
}

tunnel_example() {
	echo Starting tunnel from bastion to database...
	ssh ec2-user@${BASTION_HOST} -L 3306:${DB_HOST}:3306 -fN
	echo Verifying connection to local tunnel ...
	mysql -u ${DB_USERNAME} -p${DB_PASSWORD} -h 127.0.0.1 -e "show databases"
	killall ssh
}

load() {
	echo Creating tables and load sample CSV data into database ...
	scp ./quotes.ddl.sql ec2-user@${BASTION_HOST}:/tmp
	scp ./quotes.csv ec2-user@${BASTION_HOST}:/tmp
	ssh ec2-user@${BASTION_HOST} "mysql -u admin -p${DB_PASSWORD} -h ${DB_HOST} < /tmp/quotes.ddl.sql"
}

config-db-secret
load
ssh ec2-user@${BASTION_HOST} "mysql -u admin -p${DB_PASSWORD} -h ${DB_HOST} -e 'select count(1) from demo.quotes;'"
