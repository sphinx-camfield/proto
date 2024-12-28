protoc:
	@protoc \
		-I ./definitions \
		--go_out=pkg --go_opt=paths=source_relative $(target) \
		--go-grpc_out=pkg --go-grpc_opt=paths=source_relative $(target)
