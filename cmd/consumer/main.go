package main

import (
	"log"
	"time"

	"github.com/nats-io/go-nats-streaming"
)

func main() {
	//stan.Connect(clusterID, clientID, ops ...Option)
	ns, _ := stan.Connect("test-cluster", "consumer1", stan.NatsURL("nats://localhost:4222"))
	defer ns.Close()

	// Simple Async Subscriber
	sub, _ := ns.Subscribe("foo", func(m *stan.Msg) {
		m.Ack()
		log.Printf("Received a message: %s and reply: %S\n ", string(m.Data), m.Reply)
		//ns.Publish(m.Reply, []byte("reply.."))
	}, stan.SetManualAckMode(), stan.DurableName("i-will-remember"))
	log.Printf("subscribing to subject 'foo' \n")

	// nc := ns.NatsConn()
	// i := 3
	// nc.Subscribe("foo", func(msg *nats.Msg) {
	// 	i++
	// 	log.Printf("reply msg")
	// 	nc.Publish(msg.Reply, []byte("replya"))
	// })
	// nc.Flush()

	// msg, err := nc.Request("foo", []byte("hello jason"), 30*time.Second)
	// if err != nil {
	// 	if nc.LastError() != nil {
	// 		log.Fatalf("Error in Request1: %v\n", nc.LastError())
	// 	}
	// 	log.Fatalf("Error in Request2: %v\n", err)
	// }

	// log.Printf("Received [%v] : '%s'\n", msg.Subject, string(msg.Data))

	time.Sleep(400 * time.Second)
	// Unsubscribe
	sub.Unsubscribe()

}
