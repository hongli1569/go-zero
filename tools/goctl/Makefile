# 定义变量
BINARY_NAME=goctlX.exe
INSTALL_PATH=$(GOBIN)

# 如果GOBIN未设置，默认使用$GOPATH/bin或$HOME/go/bin
ifndef GOBIN
    ifndef GOPATH
        INSTALL_PATH=$(HOME)/go/bin
    else
        INSTALL_PATH=$(GOPATH)/bin
    endif
endif

build:
	go build -o $(BINARY_NAME) goctl.go;
install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)"
	@mkdir -p $(INSTALL_PATH)
	@mv $(BINARY_NAME) $(INSTALL_PATH)/$(BINARY_NAME)
mac:
	GOOS=darwin go build -ldflags="-s -w" -o goctl-darwin goctl.go
	$(if $(shell command -v upx || which upx), upx goctl-darwin)

win:
	GOOS=windows go build -ldflags="-s -w" -o goctl.exe goctl.go
	$(if $(shell command -v upx || which upx), upx goctl.exe)

linux:
	GOOS=linux go build -ldflags="-s -w" -o goctl-linux goctl.go
	$(if $(shell command -v upx || which upx), upx goctl-linux)

image:
	docker build --rm --platform linux/amd64 -t kevinwan/goctl:$(version) .
	docker tag kevinwan/goctl:$(version) kevinwan/goctl:latest
	docker push kevinwan/goctl:$(version)
	docker push kevinwan/goctl:latest
	docker build --rm --platform linux/arm64 -t kevinwan/goctl:$(version)-arm64 .
	docker tag kevinwan/goctl:$(version)-arm64 kevinwan/goctl:latest-arm64
	docker push kevinwan/goctl:$(version)-arm64
	docker push kevinwan/goctl:latest-arm64
