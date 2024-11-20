# Multiverse (WARNING, this is a work in progress)

## Introduction

Multiverse is a multi node Multipass cluster management system. It is designed to manage a cluster of Multipass instances.
It is aim to provide a simple way to create, start, stop, delete, shell and manage Multipass instances from any machine.
It is looks like a kubectl but for Multipass.

## Test

 > make init
 > make build
 > make test

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