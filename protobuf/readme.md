protoc -I protobuf --go_out=./generated --go-grpc_out=./generated `
          protobuf/common/*.proto `
          protobuf/auth/*.proto `
          protobuf/auth/methods/*.proto `
          protobuf/user/*.proto `
          protobuf/user/messages/*.proto `
          protobuf/user/methods/*.proto `

在根目录下执行

protoc -I protobuf `
  --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts `
  --ts_out=./generated `
  protobuf/common/*.proto `
  protobuf/auth/*.proto `
  protobuf/auth/methods/*.proto `
  protobuf/user/*.proto `
  protobuf/user/messages/*.proto `
  protobuf/user/methods/*.proto