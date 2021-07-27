# Fortune Application Docker container
# Requires temporary AWS creds in environment variables
include ../../config/config.gitignore
export AWS_REGION:=ap-southeast-1
export DB_SECRET_NAME:=service/db/mysql
export DB_HOST:=host.docker.internal
export ES_HOST:=https://host.docker.internal:35000
export REDIS_HOST:=host.docker.internal
export CONFIG_DIR=./config.gitignore
export CONFIG_FILE=${CONFIG_DIR}/fortune-config.gitignore

.DEFAULT_GOAL := help

help:
	@echo "makefile for fortune container"
	@echo "usage:"
	@echo "  make [ ... ]"

# BUILD

build-docker:
	cd .. && docker build -t fortune .

# TEST
test: stop-docker start-docker test-mysql test-es test-redis
	echo "service test completed."

test-mysql: tunnel-db
	while ! echo exit | nc -vz localhost 8080; do sleep 2; done
	-curl http://localhost:8080/fortune/sql
	-killall ssh

test-es: tunnel-es
	while ! echo exit | nc -vz localhost 8080; do sleep 2; done
	-curl http://localhost:8080/fortune/es
	-killall ssh

test-redis: tunnel-redis
	while ! echo exit | nc -vz localhost 8080; do sleep 2; done
	-curl http://localhost:8080/fortune/redis
	-killall ssh

tunnel-db:
	cd ../../fortune-data/mysql && ./tunnel-db.sh

tunnel-es:
	cd ../../fortune-data/elasticsearch && ./tunnel-es.sh

tunnel-redis:
	cd ../../fortune-data/redis && ./tunnel-redis.sh

start-docker:
	docker run -dp 8080:80 \
		-e DB_SECRET_NAME -e DB_HOST -e ES_HOST -e REDIS_HOST \
		-e AWS_REGION -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY -e AWS_SESSION_TOKEN \
		--name fortune fortune:latest 

stop-docker:
	-docker stop fortune
	-docker container rm fortune
	-killall ssh

start-kube-basic:
	kubectl apply -f fortune-single.yaml
	kubectl get deployment fortune
	minikube service fortune

stop-kube-basic:
	kubectl delete service fortune
	kubectl delete deployment fortune

set-config-minikube:
	@mkdir -p ./config.gitignore
	@rm -rf ${CONFIG_FILE}
	@echo DB_SECRET_NAME=${DB_SECRET_NAME} >> ${CONFIG_FILE}
	@echo DB_HOST=host.minikube.internal >> ${CONFIG_FILE}
	@echo ES_HOST=https://host.minikube.internal:35000 >> ${CONFIG_FILE}
	@echo REDIS_HOST=host.minikube.internal >> ${CONFIG_FILE}
	@echo AWS_REGION=${AWS_REGION} >> ${CONFIG_FILE}
	@echo AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} >> ${CONFIG_FILE}
	@echo AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} >> ${CONFIG_FILE}
	@echo AWS_SESSION_TOKEN=${AWS_SESSION_TOKEN} >> ${CONFIG_FILE}
	kubectl create configmap fortune-config --from-env-file=${CONFIG_FILE}

start-kube-config:
	kubectl apply -f fortune-config.yaml
	kubectl get deployment fortune
	minikube service fortune