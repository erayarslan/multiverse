module multipass-cluster

godebug x509negativeserial=1

go 1.23

require (
	github.com/google/uuid v1.6.0
	golang.org/x/crypto v0.29.0
	google.golang.org/grpc v1.68.0
	google.golang.org/protobuf v1.35.2
)

require (
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/term v0.26.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241007155032-5fefd90f89a9 // indirect
)
