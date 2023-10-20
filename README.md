## Install
```sh
# install protoc >= 3.15 version
# https://github.com/protocolbuffers/protobuf/releases/tag/v3.15.0
PB_REL=https://github.com/protocolbuffers/protobuf/releases
curl -LO $PB_REL/download/v3.15.0/protoc-3.15.0-linux-x86_64.zip
unzip protoc-3.15.0-linux-x86_64.zip -d $HOME/.local
echo "export PATH=\$PATH:\$HOME/.local/bin" >> $HOME/.bashrc
source ~/.bashrc
```
```sh

$ protoc \
--plugin=protoc-gen-ts_proto=./node_modules/.bin/protoc-gen-ts_proto \
--ts_proto_out=. \
--ts_proto_opt=outputServices=grpc-js,env=node,esModuleInterop=true \
./protos/auth.proto

# server
npx ts-node server.ts

# client
npx ts-node client.ts
```