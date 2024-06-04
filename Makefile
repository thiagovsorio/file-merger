install:
	- go build -o file-merger

install.go:
	rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz