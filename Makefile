include .env
export

dep:
	GOPRIVATE=go.bukalapak.io go mod download

bin:
	@mkdir -p bin

# Migration Related Commands
# https://github.com/golang-migrate/migrate
bin/migrate: bin
ifeq ($(shell uname), Linux)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar zxf - --directory /tmp && cp /tmp/migrate bin/
else
	@echo "Your OS is not supported."
endif

# If the first argument is "migrate"...
ifeq (migrate,$(firstword $(MAKECMDGOALS)))
    # use the rest as arguments for "migrate"
    MIGRATE_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
    # ...and turn them into do-nothing targets
    $(eval $(MIGRATE_ARGS):;@:)
endif

migrate: bin/migrate
	./bin/migrate -source file://db/migrations -database "mysql://$(MYSQL_USERNAME):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)" $(MIGRATE_ARGS)

#####################
# General Usage
#####################

run-backend: # generate all go generate command inside backend package.
	@go run -v app/backend/main.go

docker-start:
	@docker-compose up

docker-remove:
	@docker-compose down

