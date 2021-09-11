package main

import (
	"fmt"
	"noti/que"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	fmt.Println("Application Runing...")

	que.ConType = "tcp"
	que.Host = "localhost:9092"
	que.Partition = 0

	que.Connect()
	defer que.DisConnect()

	m := que.ListAll()

	fmt.Print(m)

	go que.Listen()

	time.Sleep(100 * time.Second)

	msg := kafka.Message{
		Key:   []byte("Github:"),
		Value: []byte("4nkitd"),
	}

	que.Speak(msg)

}
