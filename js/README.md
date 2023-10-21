## Generate
```sh
$ protoc \
--plugin=protoc-gen-ts_proto=./node_modules/.bin/protoc-gen-ts_proto \
--ts_proto_out=. \
--ts_proto_opt=outputServices=grpc-js,env=node,esModuleInterop=true \
./protos/auth.proto
```

## Start
```sh
# server
npx ts-node server.ts

# client
npx ts-node client.ts
```