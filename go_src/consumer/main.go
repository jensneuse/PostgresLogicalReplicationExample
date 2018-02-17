package main

import (
	"github.com/nats-io/go-nats-streaming"
	"log"
	"github.com/nats-io/go-nats"
	"github.com/jnsone11/PostgresLogicalReplicationTest/go_src"
	"github.com/jnsone11/PostgresLogicalReplicationTest/proto"
	"github.com/golang/protobuf/proto"
)

func main() {
	natsConn, err := nats.Connect("nats://0.0.0.0:4222")
	if err != nil {
		log.Fatal(err)
	}

	streamingConn, err := stan.Connect("test-cluster", "consumer", stan.NatsConn(natsConn))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("nats connected")

	_,err = streamingConn.QueueSubscribe(common.TOPIC, "queuegroup1", messageHandler)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("subscribed")

	select {}
}

func messageHandler(message *stan.Msg) {

	var msg decoderbufs.RowMessage
	err := proto.Unmarshal(message.Data, &msg)
	if err != nil {
		log.Println("Err: %s\n",err.Error())
		return
	}

	log.Println("--- MSG ---")
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
