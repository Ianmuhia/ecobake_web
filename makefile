DB_URL=postgresql://tvtn:tvtn@localhost:5432/rvtn?sslmode=disable

build:
	docker-compose build

ggen:
	go run github.com/99designs/gqlgen generate
geninit:
	go run github.com/99designs/gqlgen init

up:
	docker-compose up

stop:
	docker-compose down

dev:
	go run ./cmd/ *.go

db:
	mock c -f mock.yaml -r 10

reset:
	docker-compose down --remove-orphans --volumes 

shell:
	docker exec -it wyce-dev bash && cd

migrations:
	 @read -p "Migration Name? : " Name; \
 	 migrate create -ext sql -dir internal/postgresql/migrations -seq $${Name}

profile:
	 @read -p "Server Port to profile ? : " Port; \
	go tool pprof -http=:5001 $${Port}/debug/pprof/profile


heap:
	 @read -p "Server Port to profile ? : " Port; \
	go tool pprof -http=:5001 $${Port}/debug/pprof/heap


#mock:
#	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store
#
#.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 db_docs db_schema sqlc test server mock

migrateup:
	migrate -path internal/postgresql/migrations -database "$(DB_URL)" -verbose up

migrateupF:
	 @read -p "Version to force to ? : " version; \
	migrate -path internal/postgresql/migrations -database "$(DB_URL)" force $${version}

migrateDA:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrateDV:
	 @read -p "Version to downgrade to ? : " version; \
	migrate -path internal/postgresql/migrations -database "$(DB_URL)" down $${version}

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml


pgip:
	docker inspect postgres_wyce | grep IPAddress

pgdrop:
	docker stop postgres_wyce && docker rm postgres_wyce

sqlc:
	sqlc generate

#test:
#	go test -v -cover ./...

fmt:
	go fmt ./...


server:
	air

GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
BINARY_NAME=example
VERSION?=0.0.0
SERVICE_PORT?=3000
DOCKER_REGISTRY?= #if set it should finished by /
EXPORT_RESULT?=false # for CI please set EXPORT_RESULT to true

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test build vendor

all: help

## Build:
build: ## Build your project and put the output binary in out/bin/
	mkdir -p out/bin
	GO111MODULE=on $(GOCMD) build -mod vendor -o out/bin/$(BINARY_NAME) .

clean: ## Remove build related file
	rm -fr ./bin
	rm -fr ./out
	rm -f ./junit-report.xml checkstyle-report.xml ./coverage.xml ./profile.cov yamllint-checkstyle.xml

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	$(GOCMD) mod vendor

watch: ## Run the code with cosmtrek/air to have automatic reload on changes
	$(eval PACKAGE_NAME=$(shell head -n 1 go.mod | cut -d ' ' -f2))
	docker run -it --rm -w /go/src/$(PACKAGE_NAME) -v $(shell pwd):/go/src/$(PACKAGE_NAME) -p $(SERVICE_PORT):$(SERVICE_PORT) cosmtrek/air

## Test:
test: ## Run the tests of the project
ifeq ($(EXPORT_RESULT), true)
	GO111MODULE=off go get -u github.com/jstemmer/go-junit-report
	$(eval OUTPUT_OPTIONS = | tee /dev/tty | go-junit-report -set-exit-code > junit-report.xml)
endif
	$(GOTEST) -v -race ./... $(OUTPUT_OPTIONS)

coverage: ## Run the tests of the project and export the coverage
	$(GOTEST) -cover -covermode=count -coverprofile=profile.cov ./...
	$(GOCMD) tool cover -func profile.cov
ifeq ($(EXPORT_RESULT), true)
	GO111MODULE=off go get -u github.com/AlekSi/gocov-xml
	GO111MODULE=off go get -u github.com/axw/gocov/gocov
	gocov convert profile.cov | gocov-xml > coverage.xml
endif

## Lint:
lint: lint-go lint-dockerfile lint-yaml ## Run all available linters

lint-dockerfile: ## Lint your Dockerfile
# If dockerfile is present we lint it.
ifeq ($(shell test -e ./Dockerfile && echo -n yes),yes)
	$(eval CONFIG_OPTION = $(shell [ -e $(shell pwd)/.hadolint.yaml ] && echo "-v $(shell pwd)/.hadolint.yaml:/root/.config/hadolint.yaml" || echo "" ))
	$(eval OUTPUT_OPTIONS = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "--format checkstyle" || echo "" ))
	$(eval OUTPUT_FILE = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "| tee /dev/tty > checkstyle-report.xml" || echo "" ))
	docker run --rm -i $(CONFIG_OPTION) hadolint/hadolint hadolint $(OUTPUT_OPTIONS) - < ./Dockerfile $(OUTPUT_FILE)
endif

lint-go: ## Use golintci-lint on your project
	$(eval OUTPUT_OPTIONS = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "--out-format checkstyle ./... | tee /dev/tty > checkstyle-report.xml" || echo "" ))
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest-alpine golangci-lint run --deadline=65s $(OUTPUT_OPTIONS)

lint-yaml: ## Use yamllint on the yaml file of your projects
ifeq ($(EXPORT_RESULT), true)
	GO111MODULE=off go get -u github.com/thomaspoignant/yamllint-checkstyle
	$(eval OUTPUT_OPTIONS = | tee /dev/tty | yamllint-checkstyle > yamllint-checkstyle.xml)
endif
	docker run --rm -it -v $(shell pwd):/data cytopia/yamllint -f parsable $(shell git ls-files '*.yml' '*.yaml') $(OUTPUT_OPTIONS)

## Docker:
docker-build: ## Use the dockerfile to build the container
	docker build --rm --tag $(BINARY_NAME) .

docker-release: ## Release the container with tag latest and version
	docker tag $(BINARY_NAME) $(DOCKER_REGISTRY)$(BINARY_NAME):latest
	docker tag $(BINARY_NAME) $(DOCKER_REGISTRY)$(BINARY_NAME):$(VERSION)
	# Push the docker images
	docker push $(DOCKER_REGISTRY)$(BINARY_NAME):latest
	docker push $(DOCKER_REGISTRY)$(BINARY_NAME):$(VERSION)

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
