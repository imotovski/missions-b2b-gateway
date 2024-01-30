generate-gateway:
	protoc -I . --grpc-gateway_out ./gen/missions --grpc-gateway_opt paths=source_relative ./missions_service.proto

generate-server:
	protoc -I . --go_out ./gen/missions --go_opt paths=source_relative --go-grpc_out ./gen/missions/ --go-grpc_opt paths=source_relative missions_service.proto

swagger:
	protoc -I . --openapiv2_out ./gen/swagger/missions missions_service.proto
