protoc-gen:
	cd src/api && \
	protoc \
		--go_out=../pkg/grpc --go_opt=paths=source_relative \
		--go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
		$(word 2, $(MAKECMDGOALS)).proto

init:
	cd src && go mod download
	$(MAKE) protoc-gen

%:
	@: