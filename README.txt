brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc --go_out=. --go-grpc_out=. */*.proto

go build cmd/main.go
./main -master -worker
./main -worker
./main -client

                                                         ┌───────────────────────────────────┐
                                                         │       multipass client cli        │
                                                         │                                   │
                                                         │ ┌───────────────────────────────┐ │
                                                         │ │          api client           │─┼┐
 ┌───────────────────────────────────┐                   │ └───────────────────────────────┘ ││
 │   multipass cluster worker node   │                   └───────────────────────────────────┘│
 │   ┌──────┐  ┌──────┐  ┌──────┐    │                                                        │
 │   │ vm01 │  │ vm02 │  │ vm03 │    │                                                      user
 │   └──────┘  └──────┘  └──────┘    │                                                    request
 │       │         │         │       │                 ┌───────────────────────────────────┐  │
 │       └─────────┼─────────┘       │                 │   multipass cluster master node   │  │
 │                 ▼                 │                 │                                   │  │
 │      ┌────────────────────┐       │                 │ ┌───────────────────────────────┐ │  │
 │      │ multipass service  │       │                 │ │          api server           │◀┼──┘
 │      └────────────────────┘       │                 │ └───────────────────────────────┘ │
 │                 ▲                 │                 │                 │                 │
 │                 │                 │                 │                 ▼                 │
 │ ┌───────────────────────────────┐ │                 │ ┌───────────────────────────────┐ │
 │ │         agent server          │◀┼──aggregate──────┼─│         agent client          │ │
 │ └───────────────────────────────┘ │                 │ └───────────────────────────────┘ │
 │                                   │                 │                 ▲                 │
 │ ┌───────────────────────────────┐ │                 │                 │                 │
 │ │        cluster client         │─┼────────┐        │              create               │
 │ └───────────────────────────────┘ │   join with     │              agent                │
 │                                   │   agent info    │              client               │
 └───────────────────────────────────┘        │        │                 │                 │
                                              │        │                 │                 │
                                              │        │ ┌───────────────────────────────┐ │
                                              └────────┼▶│        cluster server         │ │
                                                       │ └───────────────────────────────┘ │
                                                       └───────────────────────────────────┘