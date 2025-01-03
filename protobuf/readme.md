protoc -I protobuf --go_out=./generated --go-grpc_out=./generated `
          protobuf/google/api/*.proto `
          protobuf/common/*.proto `
          protobuf/aggregator/*.proto `
          protobuf/aggregator/methods/*.proto `
          protobuf/auth/*.proto `
          protobuf/auth/messages/*.proto `
          protobuf/auth/methods/*.proto `
          protobuf/user/*.proto `
          protobuf/user/messages/*.proto `
          protobuf/user/methods/*.proto `
          protobuf/creation/*.proto `
          protobuf/creation/messages/*.proto `
          protobuf/creation/methods/*.proto

          # 根目录下的环境
protoc -I protobuf `
       --include_imports --include_source_info `
       --descriptor_set_out=./deploy/docker/envoy/descriptor.pb `
          protobuf/google/api/*.proto `
          protobuf/common/*.proto `
          protobuf/aggregator/*.proto `
          protobuf/aggregator/methods/*.proto `
          protobuf/auth/*.proto `
          protobuf/auth/messages/*.proto `
          protobuf/auth/methods/*.proto `
          protobuf/user/*.proto `
          protobuf/user/messages/*.proto `
          protobuf/user/methods/*.proto `
          protobuf/creation/*.proto `
          protobuf/creation/messages/*.proto `
          protobuf/creation/methods/*.proto

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