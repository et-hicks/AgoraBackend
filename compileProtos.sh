#!/bin/zsh

echo "proto system paths are:"
echo "\t" $AGORA_PROTO_DIR
echo "\t" $AGORA_PROTO_DST_DIR

protoc -I=$AGORA_PROTO_DIR --go_out=$AGORA_PROTO_DST_DIR $AGORA_PROTO_DIR/*.proto

go mod tidy

# run tests
# go test ./...

# go list # broken for this repo ??

echo "Done!"