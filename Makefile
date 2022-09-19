run-proto:
	protoc -I=proto --go_out=ganerete-file --go-grpc_out=ganerete-file proto/addressbook.proto