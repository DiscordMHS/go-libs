module github.com/DiscordMHS/go-libs

go 1.24.4

require (
	github.com/DiscordMHS/protocols/gen/go v0.0.24
	github.com/golang-jwt/jwt/v5 v5.3.0
	google.golang.org/grpc v1.76.0
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.0.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.3 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.29.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20251014184007-4626949a642f // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251007200510-49b9836ed3ff // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)

replace buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go => github.com/DiscordMHS/protocols/gen/go/local/buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v0.0.1
