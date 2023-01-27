# The linters to use
LINT = golangci-lint
# Flags to pass to the linter
LINT_FLAGS = run

# Build the application
build:
	go build  -o bin/main cmd/svr/main.go

# Run the application
run: build
	./bin/main

# Run the linter
lint:
	$(LINT) $(LINT_FLAGS)

# Clean up the project
clean:
	go clean
	rm -f main

# Run air (make sure you have air installed)
air:
	air

# Build and run the application in a Docker container
docker:
	docker build -t my-app .
	docker run -it --rm my-app

# Run tests
test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic

# Run benchmarks
bench:
	go test -bench=. -benchmem

# Profile the application
profile:
	go test -run=^a -bench=. -cpuprofile=cpu.out -memprofile=mem.out

swagger:
	go run github.com/go-swagger/go-swagger/cmd/swagger generate server -f swagger.yml

# Declare these as phony targets
.PHONY: build run lint clean air docker test bench profile
