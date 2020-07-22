#!/bin/bash
source ../../config/config.gitignore

tunnel-es() {
	lookup-config-db
	echo "Starting tunnel thru bastion to elasticsearch..."
	ssh ec2-user@${BASTION_HOST} -L 35000:${ES_HOST}:443 -fN
}

verify-es() {
    echo "Verifying elasticsearch connection via local tunnel ..."
    curl -k https://localhost:35000
}

tunnel-es
verify-es