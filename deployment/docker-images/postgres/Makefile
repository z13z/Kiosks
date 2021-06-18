default: build

IMAGE_NAME=zaza13/kiosks/postgres
IMAGE_VERSION=1.1

build:
	docker build -f Dockerfile . -t $(IMAGE_NAME):$(IMAGE_VERSION)

push:
	docker push $(IMAGE_NAME):$(IMAGE_VERSION)