
BASE_UBUNTU_IMAGE_NAME=ubuntu-20.04.1-live-server-amd64.iso
BASE_UBUNTU_IMAGE_VERSION=20.04.1

build: images/${BASE_UBUNTU_IMAGE_NAME} chroot_commands create_custom_image prepare_kiosk
				@echo building kiosk image from ${BASE_UBUNTU_IMAGE_NAME}
				echo ./create_custom_image ${BASE_UBUNTU_IMAGE_NAME} | sudo su

clean:
	sudo rm -rf build

images/${BASE_UBUNTU_IMAGE_NAME}:
	mkdir -p images
	curl "https://releases.ubuntu.com/${BASE_UBUNTU_IMAGE_VERSION}/${BASE_UBUNTU_IMAGE_NAME}" > $@

.PHONY: buildo
