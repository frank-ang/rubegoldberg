# Fortune Application
include ../config/config.gitignore
export AWS_REGION:=ap-southeast-1
export DB_SECRET_NAME:=service/db/mysql
export DB_HOST:=127.0.0.1
export ES_HOST:=https://127.0.0.1:35000
export REDIS_HOST:=127.0.0.1

.DEFAULT_GOAL := help

help:
	@echo "makefile for fortune"
	@echo "usage:"
	@echo "  make [ build-linux | build-mac | test | deploy-push | deploy ]"

# BUILD

build-linux:
	env GOOS=linux go build -o fortune-linux

build-mac:
	go build

# TEST
start-xray: 
	-killall xray_mac
	cd ~/bin/aws-xray-daemon-macos-3.x/ && ./xray_mac -o &

test: test-mysql test-es test-redis
	echo "service test completed."

test-connect: test-connect-mysql test-connect-es test-connect-redis
	echo "connectivity tests completed."

test-health: start-xray
	-./fortune &
	while ! echo exit | nc -vz localhost 80; do sleep 2; done
	-curl localhost
	-killall fortune

test-mysql: tunnel-db
	-./fortune -mysql &
	while ! echo exit | nc -vz localhost 80; do sleep 2; done
	#-curl localhost
	-curl localhost/fortune/sql
	-killall fortune
	-killall ssh

test-connect-mysql: tunnel-db
	-go test -v --run TestMySql fortune/mysql
	-killall ssh

test-es: tunnel-es
	-./fortune -es &
	while ! echo exit | nc -vz localhost 80; do sleep 2; done
	#-curl localhost
	-curl localhost/fortune/es
	-killall fortune
	-killall ssh

test-connect-es: start-xray tunnel-es
	-go test -v fortune/elasticsearch
	-killall ssh

test-redis: start-xray tunnel-redis
	-./fortune -redis &
	while ! echo exit | nc -vz localhost 80; do sleep 2; done
	# -curl localhost/health
	-curl localhost/fortune/redis
	-killall fortune
	-killall ssh

test-connect-redis: tunnel-redis
	-go test -v fortune/redis
	-killall ssh

tunnel-db:
	cd ../fortune-data/mysql && ./tunnel-db.sh

tunnel-es:
	cd ../fortune-data/elasticsearch && ./tunnel-es.sh

tunnel-redis:
	cd ../fortune-data/redis && ./tunnel-redis.sh

# DEPLOY

deploy-push: build-linux
	aws deploy push \
		--application-name FortuneApp \
		--s3-location s3://sandbox00-playground/codedeploy/FortuneApp.zip \
		--ignore-hidden-files

deploy: deploy-push
	aws deploy create-deployment \
		--application-name FortuneApp \
		--s3-location bucket=sandbox00-playground,key=codedeploy/FortuneApp.zip,bundleType=zip \
		--deployment-group-name FortuneAppDeploymentGroup \
		--deployment-config-name CodeDeployDefault.OneAtATime \
		--description "First Deploy"


clean: 
	go clean
