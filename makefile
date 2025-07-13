user-proto:
	rm -f grpc/user/pb/*.go
	protoc --proto_path=grpc/user/proto --go_out=grpc/user/pb --go_opt=paths=source_relative \
	--go-grpc_out=grpc/user/pb --go-grpc_opt=paths=source_relative grpc/user/proto/*.proto