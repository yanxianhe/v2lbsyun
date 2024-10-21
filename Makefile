# 定义编译时的变量
APP_NAME = v2lbsyun
BIN_DIR = bin
BUILD_TIME = $(shell date +%Y%m%d-%H%M%S)
#GIT_HASH = $(shell git rev-parse --short HEAD)
GIT_HASH = 1.2.0
LDFLAGS = -ldflags "-s -w -X main.version=$(GIT_HASH) -X main.buildTime=$(BUILD_TIME)"

# 默认目标
default: build
initenv:
	@export GO111MODULE=on
	@export GOROOT=/opt/go
	@export GOPATH=/opt/go/.mygolang
	@export GOPROXY=https://mirrors.aliyun.com/goproxy/
	@export GOPROXY=https://goproxy.io,direct
	@export PATH=$PATH:$GOROOT/bin
	@export CONTAINERS_USERNS_MODE=auto
# 编译应用
build:
	@echo "Building $(APP_NAME)..."
	@go build $(LDFLAGS) -o $(BIN_DIR)/$(APP_NAME) main.go
	@rsync -avz config.toml $(BIN_DIR)/
	@rsync -avz LICENSE $(BIN_DIR)/

# 运行应用
run:
	@echo "Running $(APP_NAME)..."
	@go run main.go

# 清理构建的二进制文件
clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)

# 显示版本信息
version:
	@echo "App Version: $(GIT_HASH)"
	@echo "Build Time: $(BUILD_TIME)"

.PHONY: build run clean version