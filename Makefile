# workdir info
PACKAGE=JiuJia
PREFIX=$(shell pwd)
CMD_PACKAGE=${PACKAGE}
OUTPUT_DIR=${PREFIX}/bin
OUTPUT_FILE=${OUTPUT_DIR}/jiujia

lint:
	@echo "+ $@"
	@$(if $(GOLINT), , \
		$(error Please install golint: `go get -u github.com/golangci/golangci-lint/cmd/golangci-lint`))
	golangci-lint run --deadline=10m ./...

test:
	@echo "+ test"
	go test -cover  ./...

build:
	@echo "+ build"
	go build -o ${OUTPUT_DIR}/jiujia ./cmd/start_without_ui

build-with-ui:
	@echo "+ build"
	go build -o ${OUTPUT_DIR}/jiujia-with-ui ./cmd/start_with_ui

clean:
	@echo "+ $@"
	@rm -r "${OUTPUT_DIR}"
