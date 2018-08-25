package main

import (
	"github.com/nats-io/go-nats-streaming"
)

func main() {

	//stan.Connect(clusterID, clientID, ops ...Option)
	ns, _ := stan.Connect("test-cluster", "publisher1", stan.NatsURL("nats://localhost:4222"))
	defer ns.Close()

	ns.Publish("foo", []byte("Hello consumer4"))

}
