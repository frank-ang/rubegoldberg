# Fortune Application
include ../config/config.gitignore
export AWS_REGION:=ap-southeast-1
export DB_SECRET_NAME:=service/db/mysql
export DB_HOST:=127.0.0.1
export ES_HOST:=https://127.0.0.1:35000
export REDIS_HOST:=redis00.v8jh5m.ng.0001.apse1.cache.amazonaws.com

.DEFAULT_GOAL := help

help:
	@echo "make it"

build:
	go build

test-mysql: tunnel-db
	-./fortune -mysql &
	while ! echo exit | nc -vz localhost 80; do sleep 2; done
	-curl localhost
	-curl localhost/fortune/sql
	-killall fortune
	-killall ssh

test-mysql-unit: tunnel-db
	-go test -v --run TestMySqlGetSecrets fortune/mysql
	-killall ssh

test-es: tunnel-es
	-./fortune -es &
	while ! echo exit | nc -vz localhost 80; do sleep 2; done
	-curl localhost
	-curl localhost/fortune/es
	-killall fortune
	-killall ssh

test-es-unit: 
	-go test -v fortune/elasticsearch

tunnel-db:
	../fortune-mysql/tunnel-db.sh

tunnel-es:
	../fortune-elasticsearch/tunnel-es.sh

clean: 
	go clean