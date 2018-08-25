package main

import (
	"log"
	"time"

	"github.com/nats-io/go-nats-streaming"
)

func main() {
	//stan.Connect(clusterID, clientID, ops ...Option)
	ns, _ := stan.Connect("test-cluster", "myID", stan.NatsURL("nats://localhost:4222"))

	// Simple Synchronous Publisher
	// does not return until an ack has been received from NATS Streaming
	// ns.Publish("foo", []byte("Hello World1"))
	// ns.Publish("foo", []byte("Hello World2"))

	// Simple Async Subscriber
	sub, _ := ns.Subscribe("foo", func(m *stan.Msg) {
		m.Ack()
		log.Printf("Received a message: %s\n", string(m.Data))
	}, stan.SetManualAckMode(), stan.DeliverAllAvailable())

	log.Printf("subscribing to subject 'foo' \n")

	time.Sleep(5 * time.Second)
	// Unsubscribe
	sub.Unsubscribe()

	// Close connection
	ns.Close()
}
