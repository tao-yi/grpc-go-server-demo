#!/bin/bash

set -e

echo "start generating pb files ... "
find . -name "*.proto" | while read -r f; do
  protoc \
    --grpc-gateway_out=. \
    --grpc-gateway_opt=logtostderr=true \
    --grpc-gateway_opt=paths=source_relative \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    --openapiv2_out=. \
    --openapiv2_opt=logtostderr=true \
    "$f"
done

echo "Done!"
