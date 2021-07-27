OUT=./pb
gen:
	protoc \
		--proto_path=./proto \
    --grpc-gateway_out=${OUT} \
    --grpc-gateway_opt=logtostderr=true \
    --grpc-gateway_opt=paths=source_relative \
    --go_out=${OUT} \
    --go_opt=paths=source_relative \
    --go-grpc_out=${OUT} \
    --go-grpc_opt=paths=source_relative \
    --openapiv2_out=./swagger \
    --openapiv2_opt=logtostderr=true \
		proto/*.proto

list:
	grpcurl -plaintext localhost:50051 list


dev:
	go run main.go