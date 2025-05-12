APP_NAME = s3-file-server
VERSION = 0.1.0
IMAGE_NAME = ghcr.io/$(GITHUB_USER)/$(APP_NAME)

build:
	go build -o bin/$(APP_NAME) .

docker-build:
	docker build --no-cache -t $(IMAGE_NAME):$(VERSION) .

docker-push: docker-build
	docker push $(IMAGE_NAME):$(VERSION)
