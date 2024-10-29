GO_MODULE := github.com/dwprz/belajar-grpc-protobuf

.PHONY: clean
clean: 
	rm --force --recursive ./protogen 
	mkdir -p ./protogen

.PHONY: protoc-go
protoc-go:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
	./proto/*.proto ./proto/type/*.proto 

.PHONY: build
build: clean protoc-go