## Generate
```sh
protoc \
--go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative \
protos/helloworld.proto

protoc \
--go_out=. \
--go-grpc_out=. \
protos/helloworld.proto
```