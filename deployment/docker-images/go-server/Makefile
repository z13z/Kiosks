default: build

IMAGE_NAME=zaza13/kiosks/master-server
IMAGE_VERSION=1.0

build:
	docker build -f Dockerfile . -t $(IMAGE_NAME):$(IMAGE_VERSION)

push:
	docker push $(IMAGE_NAME):$(IMAGE_VERSION)