
build-proto:
	protoc --proto_path=./proto \
	 	--go_out=./internal/proto \
		--go-grpc_out=./internal/proto \
		--go-grpc_opt=paths=source_relative \
		--go_opt=paths=source_relative \
			auth/auth.proto \
			message/message.proto \
			user/user.proto