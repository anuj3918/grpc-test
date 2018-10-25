# !/bin/bash

# Add the proto build scripts here for automatic building

DIR='modules'

modules=(user)

for MODULE in ${modules[@]}; do
echo "reloading $(MODULE)"
protoc -I=${DIR}/${MODULE}/proto \
	-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:${DIR}/${MODULE}/proto/ ${DIR}/${MODULE}/proto/${MODULE}.proto

protoc -I=${DIR}/${MODULE}/proto \
	-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:${DIR}/${MODULE}/proto/ ${DIR}/${MODULE}/proto/${MODULE}.proto
done