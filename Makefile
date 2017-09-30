run:
	@make docker_build
	@make docker_run
install:
	go get github.com/tools/godep; godep restore
test:
	go test -v
	go test -v ./writers
build:
	go build
docker_build:
	docker build . -t htmldrum/ds_solution
docker_run:
	docker run -it htmldrum/ds_solution
