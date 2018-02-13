package main

import (
	"github.com/jackc/pgx"
	"log"
	"context"
)

func main() {

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
	var snapshot_name string
	cp, snapshot_name, err = replicationConn.CreateReplicationSlotEx("pgx_test", "test_decoding")
	if err != nil {
		log.Fatalf("replication slot create failed: %v", err)
	}
	if cp == "" {
		log.Fatal("consistent_point is empty")
	}
	if snapshot_name == "" {
		log.Fatal("snapshot_name is empty")
	}

	err = replicationConn.StartReplication("pgx_test", 0, -1)
	if err != nil {
		log.Fatalf("Failed to start replication: %v", err)
	}

	bgContext := context.Background()

	for {
		var message *pgx.ReplicationMessage
		message, err = replicationConn.WaitForReplicationMessage(bgContext)
		if err != nil {
			log.Fatal(err)
		}

		if message.WalMessage != nil {
			log.Printf("Message:\n%s", string(message.WalMessage.WalData))
		}

		if message.ServerHeartbeat != nil {
			log.Printf("heartbeat: %s", message.ServerHeartbeat.String())
		}
	}
}
