package main

import (
	"log"
	"time"

	"github.com/nats-io/go-nats-streaming"
)

func main() {
	//stan.Connect(clusterID, clientID, ops ...Option)
	natsConn, _ := stan.Connect("test-cluster", "consumer1", stan.NatsURL("nats://10.200.252.105:4222"))
	defer natsConn.Close()

	sub := QueueSubscribe(natsConn)
	log.Printf("subscribing to subject 'delivery' \n")

	time.Sleep(500 * time.Second)
	// Unsubscribe
	sub.Unsubscribe()

}

func QueueSubscribe(natsConn stan.Conn) stan.Subscription {
	sub, _ := natsConn.QueueSubscribe("delivery", "worker1", func(m *stan.Msg) {
		m.Ack()
		log.Printf("Received a message: %s\n ", string(m.Data))
	}, stan.SetManualAckMode(), stan.DurableName("i-will-remember"))
	return sub
}
