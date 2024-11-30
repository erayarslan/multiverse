## Multiverse (this is a work in progress)

Effortless Cluster Management with Multipass

Multiverse is a robust, user-friendly tool designed to streamline the creation and management of multi-node virtual environments using Multipass. Perfect for developers and system administrators, Multiverse simplifies handling distributed setups by offering powerful features such as:

- **Cluster Creation**: Quickly spin up and configure multi-node clusters with ease.
- **Cache Synchronization**: Seamlessly sync resources between master and worker nodes for smooth operation.
- **Comprehensive CLI**: Efficiently manage nodes and instances with intuitive command-line tools.
- **Enhanced Usability**: Enjoy features like terminal resizing, shell integration, and detailed resource monitoring.

Built for scalability and flexibility, Multiverse is ideal for developing, testing, or running distributed systems in a virtualized environment. Whether you're working on a microservices architecture, experimenting with Kubernetes clusters, or running parallel workloads, Multiverse makes multi-node management straightforward and efficient.

Discover the future of simplified virtualization with Multiverse :rocket:

## Test

```shell
go run cmd/main.go --help
```

## Example Usage

```text
λ multipass start
Launched: primary
Mounted '/Users/username' into 'primary:Home'
```

```text
λ multiverse -master
master.go: master addr: localhost:1337
master.go: api server addr: localhost:1338
```

```text
λ multiverse -worker
worker.go: master to connect addr: localhost:1337
client.go: joined with uuid: $uuid
```

```text
λ multiverse -client -shell -shell-instance-name=primary
ubuntu@primary:~$
```

```text
λ multiverse -client -nodes
Node Name     IPv4                Cpu       Mem       Disk      Last Sync
hostname      127.0.0.1:*****     1         1Gb       4Gb       2024-01-01 00:00:00 UTC
```

```text
λ multiverse -client -instances
Node Name     Instance Name     State       IPv4              Image
hostname      primary           Running     xxx.xxx.xxx.xxx   ??.?? ???
```

```text
λ multiverse -client -info
Node Name     Instance Name     Cpu       Load               Disk                      Memory
hostname      primary           1         0.07 0.02 0.00     2.5GiB out of 4.0GiB      1.3GiB out of 4.0GiB
```

## Architecture
```text
                                                         ┌───────────────────────────────────┐
                                                         │       multiverse client cli       │
                                                         │                                   │
                                                         │ ┌───────────────────────────────┐ │
                                                         │ │          api client           │─┼┐
 ┌───────────────────────────────────┐                   │ └───────────────────────────────┘ ││
 │       multiverse worker node      │                   └───────────────────────────────────┘│
 │   ┌──────┐  ┌──────┐  ┌──────┐    │                                                        │
 │   │ vm01 │  │ vm02 │  │ vm03 │    │                                                      user
 │   └──────┘  └──────┘  └──────┘    │                                                    request
 │       │         │         │       │                 ┌───────────────────────────────────┐  │
 │       └─────────┼─────────┘       │                 │       multiverse master node      │  │
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
```
