# Postgres Logical Replication to Golang to Nats to everywhere..

## Quickstart

``
sh ./nats/run_nats.sh
``

``
sh ./docker/run.sh
``

``
go run ./go_src/consumer/main.go    
``

``
go run ./go_src/producer/main.go    
``

## Why?

``
Keep applications in sync with a source of truth (postgresql datbase).
``

## How?

``
With the use of postgres logical replication and the debezium decoderbufs plugin we're able to replicate database changes to a golang application via protocol buffers.
``

## Facts

- type safe
- consistent replication
- high performance
- low memory footprint