# 

## grpc 库
~~~bash
npm i -d @grpc/proto-loader
npm install ts-protoc-gen
npm install -d grpc-tools


~~~
~~~bash
#####
npm i -D ts-proto typescript
~~~
~~~sh
#!/bin/bash
protoc --plugin=node_modules/.bin/protoc-gen-ts_proto \
 --ts_proto_out=src/ \
 --ts_proto_opt=outputServices=grpc-js \
 --ts_proto_opt=esModuleInterop=true \
 -I=src/ src/proto/myapp/myapp.proto
~~~