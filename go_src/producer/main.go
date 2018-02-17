package main

import (
	"github.com/nats-io/go-nats-streaming"
	"log"
	"github.com/nats-io/go-nats"
	"github.com/jackc/pgx"
	"context"
	"github.com/jnsone11/PostgresLogicalReplicationTest/go_src"
)

const (
	TOPIC = "pgx"
)

func main() {

	natsConn, err := nats.Connect("nats://0.0.0.0:4222")
	if err != nil {
		log.Fatal(err)
	}

	streamingConn, err := stan.Connect("test-cluster", "producer", stan.NatsConn(natsConn))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("nats connected")

	startReplication(streamingConn)
}

func startReplication(streamingConn stan.Conn) {

	config := pgx.ConnConfig{
		Database: "postgres",
		Host:     "127.0.0.1",
		Port:     15432,
		User:     "pgx_replication",
		Password: "secret",
	}

	conn, err := pgx.Connect(config)
	replicationConn, err := pgx.ReplicationConnect(config)

	if err != nil {
		log.Fatal(err)
	}

	replicationConn.DropReplicationSlot("pgx_test")

	defer func() {
		replicationConn.Close()
		conn.Close()
	}()

	var cp string
	var snapshotName string
	cp, snapshotName, err = replicationConn.CreateReplicationSlotEx("pgx_test", "decoderbufs")
	if err != nil {
		log.Fatalf("replication slot create failed: %v", err)
	}
	if cp == "" {
		log.Fatal("consistent_point is empty")
	}

	if snapshotName == "" {
		log.Fatal("snapshot_name is empty")
	}

	err = replicationConn.StartReplication("pgx_test", 0, -1)
	if err != nil {
		log.Fatalf("Failed to start replication: %v", err)
	}

	log.Printf("Start replication - cp: %s; sn: %s\n", cp, snapshotName)

	bgContext := context.Background()

	for {
		var message *pgx.ReplicationMessage
		message, err = replicationConn.WaitForReplicationMessage(bgContext)
		if err != nil {
			log.Fatal(err)
		}

		if message != nil {
			if message.WalMessage != nil {

				err := streamingConn.Publish(common.TOPIC,message.WalMessage.WalData)
				if err != nil {
					log.Printf("Error Publishing: %s\n",err.Error())
				} else {
					log.Printf("Published: %d\n",message.WalMessage.WalStart)
				}

			} else if message.ServerHeartbeat != nil {
				log.Printf("heartbeat: %s", message.ServerHeartbeat.String())
			} else {
				log.Println("--- RAW ---")
			}
		}
	}
}
