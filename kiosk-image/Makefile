
BASE_UBUNTU_IMAGE_NAME=ubuntu-20.04.2-live-server-amd64.iso
BASE_UBUNTU_IMAGE_VERSION=20.04.2

build: images/${BASE_UBUNTU_IMAGE_NAME} chroot_commands create_custom_image prepare_kiosk
				@echo building kiosk image from ${BASE_UBUNTU_IMAGE_NAME}
				sudo chmod +x create_custom_image
				sudo chmod +x prepare_kiosk
				sudo chmod +x chroot_commands
				echo ./create_custom_image ${BASE_UBUNTU_IMAGE_NAME} ${KIOSK_IMAGE_ID} ${LOCAL_MACHINE} ${KIOSK_APP_NAME} "\"${CUSTOM_SCRIPTS_TO_RUN}\"" | sudo su
				rm -rf images

clean:
	sudo rm -rf build

images/${BASE_UBUNTU_IMAGE_NAME}:
	mkdir -p images
	cp ../../../*.iso $@

.PHONY: build
