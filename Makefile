install:
	go get github.com/tools/godep; godep restore
test:
	go test -v
	go test -v ./writers
build:
	go build
