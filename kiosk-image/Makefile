
BASE_UBUNTU_IMAGE_NAME=ubuntu-20.04.1-live-server-amd64.iso

build: images/${BASE_UBUNTU_IMAGE_NAME} chroot_commands create_custom_image prepare_kiosk
				@echo building kiosk image from ${BASE_UBUNTU_IMAGE_NAME}
				./create_custom_image ${BASE_UBUNTU_IMAGE_NAME}

clean:
	sudo umount -R --force build
	rm -rf build


.PHONY: buildo
