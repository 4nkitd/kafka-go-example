package que

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

var (
	Conn *kafka.Conn

	Host      string
	Topic     string   = "Notification"
	Partition int      = 0
	Brokers   []string = []string{Host}
	ConType   string

	CTX context.Context = context.Background()
)

func Connect() {

	con, err := kafka.DialLeader(
		CTX,
		ConType,
		Host,
		Topic,
		Partition,
	)

	if err != nil {
		log.Fatal("failed to dial leader:", err)
		panic(err)
	}

	Conn = con

}

func DisConnect() {

	if err := Conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
		panic(err)
	}

}
