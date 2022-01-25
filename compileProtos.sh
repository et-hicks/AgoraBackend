#!/bin/zsh

echo "proto system paths are:"
echo "\t" $AGORA_PROTO_DIR
echo "\t" $AGORA_PROTO_DST_DIR

# remove old compiled files
cd protobufs && rm -rf protobufs

protoc -I=$AGORA_PROTO_DIR --go_out=$AGORA_PROTO_DST_DIR $AGORA_PROTO_DIR/*.proto

go mod tidy # TODO: figure out how to make this tidy work

# run tests
# go test ./...

# go list # broken for this repo ??

echo "Done!"