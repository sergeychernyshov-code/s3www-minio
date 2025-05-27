APP_NAME = s3-file-server
VERSION = 0.1.0
IMAGE_NAME = ghcr.io/$(GITHUB_USER)/$(APP_NAME)

build:
	go build -o bin/$(APP_NAME) .

docker-push: 
	docker buildx build --platform linux/amd64,linux/arm64 -t $(IMAGE_NAME):$(VERSION) --push .
