#!/usr/bin/env bash

protoc -I. -I/var/www/googleapis/ --proto_path=proto/ --go_out=plugins=grpc:. proto/api.proto
protoc -I. -I/var/www/googleapis/ --grpc-gateway_out=logtostderr=true:. proto/api.proto
protoc -I. -I/var/www/googleapis/ --swagger_out=logtostderr=true:. proto/api.proto