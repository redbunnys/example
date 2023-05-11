#!/bin/bash
# Path to this plugin, Note this must be an abolsute path on Windows (see #15)
PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"
# Path to the grpc_node_plugin
PROTOC_GEN_GRPC_PATH="./node_modules/.bin/grpc_tools_node_protoc_plugin"
OUT_DIR="./src/proto/myapp"

protoc \
    # 生成消息类型 .js 和 .d.ts 文件
    --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
    # 生成 grpc 相关 .js
    --plugin="protoc-gen-grpc=${PROTOC_GEN_GRPC_PATH}" \
    --js_out="import_style=commonjs,binary:${OUT_DIR}" \
    # 传递参数, 生成 grpc .d.ts 文件, 并指明我们使用的是 @grpc/grpc-js
    --ts_out="service=grpc-node,mode=grpc-js:${OUT_DIR}" \
    --grpc_out="grpc_js:${OUT_DIR}" \
    hello.proto
