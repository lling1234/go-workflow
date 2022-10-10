#!/bin/bash
goctl rpc protoc source/rpc/kernel.proto --go_out=./kernel --go-grpc_out=./kernel/ --zrpc_out=./kernel/ --style=goZero

protoc --go_out=. --go_opt=paths=import source/rpc/message.proto

protoc-go-inject-tag -input=rpc/kernel/message.pb.go

ent message ./source/ent/schema/

goctl api plugin -p goctl-swagger="swagger -filename swagger.json" --api source/api/wflow.api --dir doc/swagger/