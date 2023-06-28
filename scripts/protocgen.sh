#!/usr/bin/env bash

set -eo pipefail

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep "option go_package" $file &> /dev/null ; then
      buf generate --template buf.gen.gogo.yml $file
    fi
  done
done

# command to generate docs using protoc-gen-doc
# buf protoc \
# -I "proto" \
# -I "third_party/proto" \
# --doc_out=./docs/core \
# --doc_opt=./docs/protodoc-markdown.tmpl,proto-docs.md \
# $(find "$(pwd)/proto" -maxdepth 5 -name '*.proto')
# go mod tidy

# # generate codec/testdata proto code
# buf protoc -I "proto" -I "third_party/proto" -I "testutil/testdata" --gocosmos_out=plugins=interfacetype+grpc,\
# Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:. ./testutil/testdata/*.proto

# move proto files to the right places
cp -r github.com/bandprotocol/bandchain-packet/* ./
rm -rf github.com
