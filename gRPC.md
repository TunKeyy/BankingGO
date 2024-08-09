## Install gRPC compiler
```
$ apt install -y protobuf-compiler
```
And, install 2 Go plugins as well:
```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## For testing api of gRPC purpose - Using gRPC client
We gonna use evans package

```
$ wget https://github.com/ktr0731/evans/releases/download/v0.10.0/evans_linux_amd64.tar.gz
$ tar -xzf evans_linux_amd64.tar.gz
$ sudo mv evans /usr/local/bin

$ evans --version
```

Connect to server
```
evans --host localhost --port 9090 -r repl
```

If it doesn't work, try using this coommand:
```
evans --proto proto/service_simple_bank.proto --host localhost --port 9090 repl
```

*Don't forget to start server first

See more about integrate grpc in https://grpc.io/docs/languages/go/quickstart/


## For gRPC gateway session - follow the following link:
https://github.com/grpc-ecosystem/grpc-gateway