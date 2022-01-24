APP = jim_service
PACKAGE =
OUTPUT_BUILD_DIR = /data

.PHONY :a help proto tidy download build run compose clean gormt gen token test

help:
	@echo "make proto -  grpc编译"
	@echo "make tidy -  Go Mod tidy"
	@echo "make download -  Go Mod下载"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make compose - docker-compose直接运行 Go 代码"
	@echo "make clean - 清除vendor"
	@echo "make gormt - 使用gormt自动生成model"
	@echo "make test - 执行测试代码"
	@echo "make grpcurl - 查看当前的服务"
	@echo "make start - goreman start"
	@echo "make status - goreman run status"
	@echo "make stop - goreman run stop"
	@echo "make proto_update -proto 更新"
	@echo "make command -显示所有bin目录命令"

proto: download
	@protoc -I=internal/jim_proto/proto \
    		-I=$(GOPATH)/pkg/mod \
    		-I=$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
    		-I=$(GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.3 \
          --doc_out=internal/jim_proto/ --doc_opt=html,docs.html  internal/jim_proto/proto/*.proto

	@protoc -I=internal/jim_proto/proto \
		-I=$(GOPATH)/pkg/mod \
		-I=$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		-I=$(GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.3 \
      --go_out=internal/jim_proto/ --go-grpc_out=internal/jim_proto/ --grpc-gateway_out=internal/jim_proto/ --validate_out="lang=go:internal/jim_proto/" \
      --doc_out=internal/jim_proto/ --doc_opt=markdown,docs.md  \
      --grpc-gateway_opt logtostderr=true internal/jim_proto/proto/*.proto

tidy:
	@go mod tidy
download:
	@go mod download
build:
	@echo "Building  app..."
	@rm -rf ${OUTPUT_BUILD_DIR}
	@mkdir -p ${OUTPUT_BUILD_DIR}
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o ${OUTPUT_BUILD_DIR}/${APP} -ldflags '-w -s'
run: build
	@${OUTPUT_BUILD_DIR}/${APP}
compose:
	@echo "Building ${APP} app in docker..."
	@docker-compose up --remove-orphans -d
clean:
	-echo "Cleaning..."
	-rm -rf vendor
gormt:
	@gormt
gen:
	@go run cmd/main/main.go gen_gorm
token:
	@go run cmd/main/main.go gen_token
test:
	@go test -v  ./test
grpcurl:
	@ grpcurl -plaintext 127.0.0.1:9093 list
grpcui:run
	@grpcui -plaintext 127.0.0.1:9093
start:build
	@goreman start
status:
	@goreman run status
stop:
	@goreman run stop
check:
proto_update:
	@cd internal/jim_proto && git pull origin master
command:
	@echo "goreman  gormt  grpcui  grpcurl  protoc-gen-doc  protoc-gen-go  protoc-gen-go-grpc  protoc-gen-govalidators  protoc-gen-grpc-gateway  protoc-gen-openapiv2  protoc-gen-validate"

