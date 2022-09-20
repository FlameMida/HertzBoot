docs:
	@echo "格式化文档......"
	swag fmt
	@echo "格式化文档完成"
	@echo "生成admin文档......"
	swag init
	@echo "生成api文档成功"
	@echo "finished"

build:
	GOAMD64=v3 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"  -o ./bin/ .
.PHONY: docs  build
