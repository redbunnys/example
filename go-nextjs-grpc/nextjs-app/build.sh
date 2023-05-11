#!/bin/bash
protoc --plugin=node_modules/.bin/protoc-gen-ts_proto \
 --ts_proto_out=src/ \
 --ts_proto_opt=outputServices=grpc-js \
 --ts_proto_opt=esModuleInterop=true \
 -I=src/ src/proto/myapp/myapp.proto