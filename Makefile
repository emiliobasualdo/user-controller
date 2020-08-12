# https://github.com/azer/go-makefile-example/blob/master/Makefile

## run: Runs application.go. ENV=DEV
run:
	ENV=DEV
	go run application.go

## watch: Auto-starts when code changes. ENV=DEV
watch:
	ENV=DEV
	@echo "  >  Dev watch mode. ¡Does NOT produce swagger documentation!"
	reflex -sr '.*.go' -G 'docs/*'  sh scripts/updateDocAndRun.sh

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
	swag init -g pkg/webapp/apiinfo.go

## test: Runs test. ENV=TEST
test:
	@echo "  >  Running Tests"
	go test ./... -v

## docker [ENV]: Builds docker image and runs the application in port 5000. Must specify ENV
docker:
	@echo "  >  Building Docker image and hosting application on port 5000"
	ENV=$(env)
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