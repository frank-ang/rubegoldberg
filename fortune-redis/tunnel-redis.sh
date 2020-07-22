#!/bin/bash
source ../config/config.gitignore

tunnel-redis() {
	lookup-config-db
	echo "Starting tunnel thru bastion to elasticsearch..."
	ssh ec2-user@${BASTION_HOST} -L 6379:${REDIS_HOST}:6379 -fN
}

verify-redis() {
    echo "Verifying redis connection via local tunnel ..."
    nc -vz 127.0.0.1 6379
}

tunnel-redis
verify-redis