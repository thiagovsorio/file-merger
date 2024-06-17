docker.build.image:
	docker build -t file-merger .

run:
	@echo ${FILE_ONE} ${FILE_TWO} ${FILE_DEST}
	docker run -it \
	-v ./:/app \
	-e FILE_ONE=${FILE_ONE} \
	-e FILE_TWO=${FILE_TWO} \
	-e FILE_DEST=${FILE_DEST} \
	 --rm file-merger
install:
	- go build -o file-merger

install.go:
	rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz