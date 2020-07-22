#!/bin/bash
source ../config/config.functions

tunnel-db() {
	lookup-config-db
	echo "Starting tunnel thru bastion to database..."
	ssh ec2-user@${BASTION_HOST} -L 3306:${DB_HOST}:3306 -fN
}

verify-db() {
    echo "Verifying database connection via local tunnel ..."
	mysql -u ${DB_USERNAME} -p${DB_PASSWORD} -h 127.0.0.1 -e "show databases"
}

tunnel-db
verify-db