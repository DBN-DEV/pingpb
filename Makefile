gen-go: install-gogrpc install-vtprotobuf gen-go-pb

gen-go-pb:
	mkdir -p ./gopb && protoc --go-grpc_out=./gopb --go_out=./gopb --go-vtproto_out=./gopb --go-vtproto_opt=features=marshal+unmarshal+size ./proto/*.proto

install-gogrpc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

install-vtprotobuf:
	go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@latest