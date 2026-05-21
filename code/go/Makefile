export WORKSPACE = $(shell pwd)
PLATFORM=$(shell uname -m)
DATETIME=$(shell date "+%Y%m%d%H%M%S")
VERSION=v2.2.1
BUILD_DIR_REST  := rest-server
BUILD_NAME_REST := rest-server-bin
release:
	@GOPATH=${GOPATH} go build  -o $(WORKSPACE)/${BUILD_NAME_REST} main.go
	@mkdir -p ./release
	@rm -rf ./tmp/rest-server/
	@mkdir -p ./tmp/rest-server/
	@mkdir ./tmp/rest-server/static
	@mkdir ./tmp/rest-server/configs
	@mkdir  ./tmp/rest-server/storage
	@mkdir  ./tmp/rest-server/views
	@cp -r  static/* ./tmp/rest-server/static
	@cp -r storage/* ./tmp/rest-server/storage
	@cp -r views/* ./tmp/rest-server/views
	@mv rest-server-bin ./tmp/rest-server/
	@cp -r configs/* ./tmp/rest-server/configs
	@cd ./tmp;tar -zcvf rest-server-$(VERSION).$(DATETIME).$(PLATFORM).tar.gz rest-server; mv  rest-server-$(VERSION).$(DATETIME).$(PLATFORM).tar.gz ../release
	@rm -rf ./tmp/
clear:
	@rm -rf  ./release
start:
	@cd bin &&  ./rest-server-bin






