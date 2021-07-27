gen:
	./protogen.sh

list:
	grpcurl -plaintext localhost:50051 list
