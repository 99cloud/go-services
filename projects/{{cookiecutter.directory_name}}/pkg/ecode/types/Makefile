PROTO_FILES=$(sort $(wildcard ./*.proto))
PROTOC_INC_PATH=/usr/local/include

#PROTOC_FLAGS:=-I. -I$(GOOGLEAPIS_PATH) -I$(GRPC_GATEWAY_PATH) -I$(PROTOC_INC_PATH) -I$(VENDOR_PATH)
PROTOC_FLAGS:=-I. -I$(PROTOC_INC_PATH)

GOPATH:=$(shell go env GOPATH)
PWD:=$(shell pwd)

generate: $(PROTO_FILES) Makefile

	protoc $(PROTOC_FLAGS) --go_out=plugins=grpc:. ${PROTO_FILES}
	@echo "ok"
