Generates Proxy stubs to proxy package

`protoc -I . --grpc-gateway_out ./gen \
--grpc-gateway_opt paths=source_relative \
./service.proto
`


Generates server stubs

`protoc -I . \
--go_out ./gen --go_opt paths=source_relative \
--go-grpc_out ./gen/ --go-grpc_opt paths=source_relative \
service.proto
`