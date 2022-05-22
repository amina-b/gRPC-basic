# This is simple example of using golang gRPC


# Install protoc compiler

First we need to instal complier for protocol buffer.
On MacOS you can use this command:
```brew install protobuf``` and for other OS you can find more about installation on this link : 
https://github.com/protocolbuffers/protobuf/releases/tag/v3.20.1 


Plugin for GoLang is on this link https://developers.google.com/protocol-buffers/docs/gotutorial 


Command ```protoc --go_out=. example.proto``` is used for generating `.pb.go file.` This file contains structs and additional methods that are generated from the example.proto file. Flag --go-out represents location where the result file will be located, and example.proto is provided file which we want to compile. 

Command  `protoc --go-grpc_out=. example.proto` is used for generating `example_grpc.pb.go` file. This file contains functions, interfaces and methods to register client and server.

Here will used these commands because of packages and folder structure.

Command for generating .pb.go file is:

```protoc -I $GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/amina-b/gRPC-basic/models/user.proto```

Command for generating _grpc.pb.go file is:

```protoc -I $GOPATH/src --go-grpc_out=require_unimplemented_servers=false:$GOPATH/src $GOPATH/src/github.com/amina-b/gRPC-basic/models/user.proto``

