主要的4个配置：
  - GOPATH ：配置GO的工作目录
  - PATH：将 $GOPATH/bin 目录加入该环境变量，可以直接执行安装的命令
  - GO111MODULE:  配置 "on" 开启 go mod
  - GOPROXY：配置 "https://goproxy.io" 开启代理

GRPC涉及的库：
  - google.golang.org/`grpc`
  - google.golang.org/protobuf
  - github.com/golang/protobuf/protoc-gen-go
  - google.golang.org/grpc/cmd/protoc-gen-go-grpc
  
编译命令：
  - protoc --proto_path=/vagrant/go_project/src/jgo/protected/proto  --go_out=plugins=grpc:/vagrant/go_project/src/jgo/protected/proto/build  *.proto
bin目录命令：
  - goreman  gormt  grpcui  grpcurl  protoc-gen-doc  protoc-gen-go  protoc-gen-go-grpc  protoc-gen-govalidators  protoc-gen-grpc-gateway  protoc-gen-openapiv2  protoc-gen-validate