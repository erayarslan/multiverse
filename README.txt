# Multiverse (WARNING, this is a work in progress)

## Introduction

Multiverse is a multi node Multipass cluster management system. It is designed to manage a cluster of Multipass instances.
It is aim to provide a simple way to create, start, stop, delete, shell and manage Multipass instances from any machine.
It is looks like a kubectl but for Multipass.

## Test

 > make init
 > make build
 > make test

## Example Usage

λ multipass start
Launched: primary
Mounted '/Users/username' into 'primary:Home'

λ multiverse -master
master.go: master addr: localhost:1337
master.go: api server addr: localhost:1338

λ multiverse -worker
worker.go: master to connect addr: localhost:1337
client.go: joined with uuid: $uuid

λ multiverse -client -shell -shell-instance-name=primary
ubuntu@primary:~$

λ multiverse -client -nodes
Node Name     IPv4                Last Sync
hostname      127.0.0.1:*****     1337-13-37 13:37:17 UTC

λ multiverse -client -instances
Node Name     Instance Name     State       IPv4              Image
hostname      primary           Running     xxx.xxx.xxx.xxx   Ubuntu ??.?? ???

## Architecture
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