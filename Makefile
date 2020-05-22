all: help

help:
	@echo "Cartesian API development"
	@echo ""
	@echo "run                         Run the application on the development environment"
	@echo "tests                       Execute the tests in the development environment"
	@echo "coverage                    Generate test coverage in the development environment"
	@echo "deps                        Update and vendorize the application dependencies"
	@echo ""

run:

tests:
	@go test ./...

coverage:
	@go test ./... -coverprofile=/tmp/coverage.out -coverpkg=./...
	@go tool cover -html=/tmp/coverage.out

deps:
	@go mod tidy
	@go mod vendor