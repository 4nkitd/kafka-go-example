package que

import (
	"log"

	"github.com/segmentio/kafka-go"
)

func Speak(msg kafka.Message) {

	_, err := Conn.WriteMessages(msg)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

}
