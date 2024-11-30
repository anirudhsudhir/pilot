run: pilot utils/sample-rootfs.tar
	sudo ./pilot run /bin/bash

pilot: main.go
	go build

utils/sample-rootfs.tar:
	mkdir -p utils/sample-rootfs/
	docker export $$(docker create ubuntu) --output="utils/sample-rootfs.tar"
	tar -xzf utils/sample-rootfs.tar -C utils/sample-rootfs
	touch utils/sample-rootfs/CONTAINER_ROOTFS

.PHONY: clean rootfs

rootfs: utils/sample-rootfs.tar

clean:
	rm -rf utils/
