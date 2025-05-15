APP_NAME := gprc_sample

.PHONY: all build clean run

all: build

build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(APP_NAME) main.go

run: build
	@echo "Running $(APP_NAME)..."
	@./$(APP_NAME)


clean: rm -f $(APP_NAME)