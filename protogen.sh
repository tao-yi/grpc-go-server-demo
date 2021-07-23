#!/bin/bash

set -e

echo "start generating pb files ... "
find . -name "*.proto" | while read -r f; do
  protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    "$f"
done

echo "Done!"
