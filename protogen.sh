#!/bin/bash

set -e

OUT="${PWD}/pb"
echo "start generating pb files ... "
find . -name "*.proto" | while read -r f; do
  protoc \
    --grpc-gateway_out="${OUT}" \
    --grpc-gateway_opt=logtostderr=true \
    --grpc-gateway_opt=paths=source_relative \
    --go_out=."${OUT}" \
    --go_opt=paths=source_relative \
    --go-grpc_out="${OUT}" \
    --go-grpc_opt=paths=source_relative \
    --openapiv2_out="${OUT}" \
    --openapiv2_opt=logtostderr=true \
    "$f"
done

echo "Done!"
