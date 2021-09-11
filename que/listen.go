package que

import (
	"fmt"
	"time"
)

func Listen() {

	for {

		msg, err := Conn.ReadMessage(10)
		if err != nil {
			fmt.Println(" Error : ", err)
		}

		fmt.Println("Msg : ", string(msg.Key), string(msg.Value), msg.Topic, msg.Partition, msg.Offset)
	}

}

func SetDeadLine(min time.Duration) {
	Conn.SetWriteDeadline(time.Now().Add(min * time.Second))
}
