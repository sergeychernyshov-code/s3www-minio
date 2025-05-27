APP_NAME = s3-file-server
VERSION := 0.1.$(shell git rev-list --count HEAD)-$(shell git rev-parse --short HEAD)
IMAGE_NAME = ghcr.io/$(GITHUB_USER)/$(APP_NAME)

build:
	go build -o bin/$(APP_NAME) .

docker-push: 
	docker buildx build --platform linux/amd64,linux/arm64 -t $(IMAGE_NAME):$(VERSION) --push .
