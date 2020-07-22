#!/bin/bash
source ../../config/config.gitignore
source ../../config/config.functions

load() {
	echo Creating tables and load sample CSV data into database ...
	scp ./quotes.ddl.sql ec2-user@${BASTION_HOST}:/tmp
	scp ./quotes.csv ec2-user@${BASTION_HOST}:/tmp
	ssh ec2-user@${BASTION_HOST} "mysql -u admin -p${DB_PASSWORD} -h ${DB_HOST} < /tmp/quotes.ddl.sql"
}

lookup-config-db
load
ssh ec2-user@${BASTION_HOST} "mysql -u admin -p${DB_PASSWORD} -h ${DB_HOST} -e 'select count(1) from demo.quotes;'"
