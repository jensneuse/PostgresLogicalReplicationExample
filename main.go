package main

import (
	"github.com/jackc/pgx"
	"log"
	"context"
	"github.com/jnsone11/PostgresLogicalReplicationTest/proto"
	"github.com/gogo/protobuf/proto"
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

	log.Printf("Start replication - cp: %s; sn: %s\n",cp, snapshotName)

	bgContext := context.Background()

	for {
		var message *pgx.ReplicationMessage
		message, err = replicationConn.WaitForReplicationMessage(bgContext)
		if err != nil {
			log.Fatal(err)
		}

		if message != nil {
			if message.WalMessage != nil {

				var msg decoderbufs.RowMessage
				err = proto.Unmarshal(message.WalMessage.WalData, &msg)
				if err != nil {
					log.Println(err.Error())
				} else {

					log.Println("--- MSG ---")

					log.Printf("WalStart: %d", message.WalMessage.WalStart)
					log.Printf("ServerWalEnd: %d", message.WalMessage.ServerWalEnd)
					log.Printf("Table: %s", msg.GetTable())
					log.Printf("Operation: %s", msg.Op.String())

					for _, field := range msg.NewTuple {
						log.Println("--- FIELD ---")

						log.Printf("Field str: %s", field.String())

						log.Printf("Field ColumnName: %s", field.GetColumnName())
						log.Printf("Field ColumnType: %d", field.GetColumnType())

						switch field.GetColumnType() {
						case int64(23):
							log.Printf("Field Int32: %d", field.GetDatumInt32())
						case int64(25):
						case int64(3802):
							log.Printf("Field String: %s", field.GetDatumString())
						}

						log.Println("--- FIELD_END ---")
					}

					log.Println("--- END ---")
				}

			} else if message.ServerHeartbeat != nil {
				log.Printf("heartbeat: %s", message.ServerHeartbeat.String())
			} else {
				log.Println("--- RAW ---")
			}
		}
	}
}
