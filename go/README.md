## Plugin
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> $HOME/.bashrc
source $HOME/.bashrc
```