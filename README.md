1. wget https://github.com/protocolbuffers/protobuf/releases/download/v23.0/pratoc-23.
   e-linux-x86 54.21p
2. unzip protoc-23.0-linux-x86 64.zip -d protec
3. sudo mv protoc/bin/protoc /usr/local/bin/
4. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
5. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
6. echo $PATH
7. export PATH="$PATH:$(go env GOPATH)/bin"
8. source ~/.bashrc
9. create the pb falder in your project
10. add this to account.proto option go package "github.com/ebhlz88/go-grpc-graphql-micro/account/pb":
11. finally run this command - protoc --go_out=./pb --go-grpc_out=./pb account.proto

-- generate graphql
go run github.com/99designs/gqlgen generate
