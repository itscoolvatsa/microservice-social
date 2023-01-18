SWAGGER_UI_VERSION:=v4.15.5

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    --openapiv2_out=pb/swagger --openapiv2_opt=allow_merge=true,merge_file_name=auth \
    proto/*.proto

generate:
   	SWAGGER_UI_VERSION=$(SWAGGER_UI_VERSION) ./scripts/generate-swagger-ui.sh

.PHONY: proto generate