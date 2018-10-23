package main

import (
	"github.com/nats-io/go-nats-streaming"
)

func main() {

	//stan.Connect(clusterID, clientID, ops ...Option)
	ns, _ := stan.Connect("test-cluster", "publisher1", stan.NatsURL("nats://10.200.252.105:4222"))
	defer ns.Close()

	ns.Publish("delivery", []byte("Hello consumer4"))

}
