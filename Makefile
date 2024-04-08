DOCKER_CONTAINER_NAME = "go_tree_test_container"

all: build

build:
	go build

test:
	go test

format:
	go fmt *.go

docker:
	docker build -t $(DOCKER_CONTAINER_NAME) .
	docker run --rm -it $(DOCKER_CONTAINER_NAME)