include .env

GO_SRC_DIR = cmd/server
BIN_DIR = bin
GO_BIN_NAME = ezd-daemon
GO_BIN_PATH = ${BIN_DIR}/${GO_BIN_NAME}

build:
	go build -o $(GO_BIN_PATH) ${GO_SRC_DIR}/main.go
run:
	./$(GO_BIN_PATH)
watch: build
	air --build.cmd "make build" --build.bin "make run"