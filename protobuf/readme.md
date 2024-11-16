protoc -I protobuf --go_out=./generated --go-grpc_out=./generated `
          protobuf/common/*.proto `
          protobuf/auth/*.proto `
          protobuf/auth/methods/*.proto `
          protobuf/user/*.proto `
          protobuf/user/messages/*.proto `
          protobuf/user/methods/*.proto `

在根目录下执行