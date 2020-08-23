# https://github.com/azer/go-makefile-example/blob/master/Makefile

## run: Runs application.go. ENV=DEV
run: export ENV = DEV
run:
	@echo "  >  Simple go run"
	go run application.go

## watch: Auto-starts when code changes. ENV = DEV
watch: export ENV = DEV
watch:
	@echo "  >  Dev watch mode. Produces swagger documentation and runs"
	reflex -sr '.*.go' -G 'docs/*' $(MAKE) documentation $(MAKE) run

## install: Equivalent to go get
install:
	@echo "  >  Downloading dependencies"
	go get

## clean: Equivalent to go clean
clean:
	@echo "  >  Cleaning build cache"
	go clean

## documentation: Generates Swagger documentation
documentation:
	@echo "  >  Generating Swagger documentation"
	swag init -g internal/webapp/apiinfo.go

## test: Runs test. ENV=TEST
test:
	@echo "  >  Running Tests"
	go test ./... -v

## docker [ENV]: Builds docker image and runs the application in port 5000. Must specify ENV
docker:
	@echo "  >  Building Docker image and hosting application on port 5000"
	ENV=$(env)
	go clean
	docker build --build-arg var_name=${ENV} --tag user-controller:latest .
	docker run --publish 5000:5000 user-controller:latest

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo