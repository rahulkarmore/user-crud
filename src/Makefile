BINARY_NAME= user-crud

# build-and-run combo (build and run at same time)
.PHONY:build-and-run
build-and-run:
	make build run-build

.PHONY: build
build:
	go build -o ${BINARY_NAME}-windows main.go
.PHONY: run-build
run-build:
	./${BINARY_NAME}-windows

# dependencies installing dependencies
.PHONY: dependencies
dependencies: 
	@echo "Installing dependencies"
	go mod tidy

# run is for run out project noramly 
.PHONY: run
run: 
	go run main.go