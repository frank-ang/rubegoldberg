# Fortune Application
include ../config/config.gitignore

.DEFAULT_GOAL := help

help:
	@echo "makefile for fortune"
	@echo "usage:"
	@echo "  make [ test... ]"

test-rest-redis:
	curl "https://$(API_GW_HOSTNAME_REST)/prod/fortune/redis"

test-rest-secure:
	@# Send request with bearer token.
	$(eval TOKEN := $(shell ./auth-cognito.sh))
	curl -H "Accept: application/json" \
		 -H 'Authorization: Bearer $(TOKEN)' \
		 "https://$(API_GW_HOSTNAME_REST)/prod/secure"

test-http-redis:
	curl "https://$(API_GW_HOSTNAME_HTTP)/fortune/redis"
	@echo

test-http-sql-auth:
	@# Send request with bearer token.
	$(eval TOKEN := $(shell ./auth-cognito.sh))
	curl -H "Accept: application/json" \
		 -H 'Authorization: Bearer $(TOKEN)' \
		 "https://$(API_GW_HOSTNAME_HTTP)/fortune/sql" 

test-http-sql:
	curl "https://$(API_GW_HOSTNAME_HTTP)/fortune/sql"
	@echo
