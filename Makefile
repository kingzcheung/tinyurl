GOCMD=GOPROXY=https://goproxy.cn GONOPROXY=*.ifreegroup.cn GONOSUMDB=*.ifreegroup.cn GOPRIVATE=*.ifreegroup.cn  /usr/bin/env go
BINARY_NAME=tinyurl
GO_MAIN=./cmd/tinyurl/main.go

.PHONY: all

default: build

build: build-web proto build-linux

build-staging: build-web-staging proto build-linux

build-web-staging:
	cd web && yarn --registry=https://registry.npm.taobao.org/ && yarn run staging
	cd web && cp -r dist ../static/

build-web:
	cd web && yarn --registry=https://registry.npm.taobao.org/ && yarn run build
	cd web && cp -r dist ../static/

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build -ldflags="-w -s" -o ./dist/bin/$(BINARY_NAME) $(GO_MAIN)


proto:
	protoc --go_out=pbgen \
		--go_opt=paths=source_relative \
		--proto_path=./pb/authcenter \
        --go-grpc_out=pbgen \
        --go-grpc_opt=paths=source_relative \
         auth.proto