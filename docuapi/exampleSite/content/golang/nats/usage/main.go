package main

import (
	"fmt"
	"log"
	"time"

	nats "github.com/nats-io/nats.go"
)

const timeout = 1 * time.Second
const DefaultURL =  "nats://39.105.141.168:4222"

func main() {
	// Connect Options.
	func setupConnOptions(opts []nats.Option) []nats.Option {
		totalWait := 10 * time.Minute
		reconnectDelay := time.Second

		opts = append(opts, nats.ReconnectWait(reconnectDelay))
		opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
		opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
		opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
		opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
		return opts
	}

	opts := []nats.Option{nats.Name("NATS Sample Publisher")}

	// Use UserCredentials
	if *userCreds != "" {
		opts = append(opts, nats.UserCredentials(*userCreds))
	}

	// Connect to NATS
	nc, err := nats.Connect(*urls, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Connect to a server

	nc, _ := nats.Connect(DefaultURL)

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Responding to a request message
	nc.Subscribe("request", func(m *nats.Msg) {
		m.Respond([]byte("answer is 42"))
	})

	// Simple Sync Subscriber
	sub, err := nc.SubscribeSync("foo")
	fmt.Println(err)
	m, err := sub.NextMsg(timeout)
	fmt.Println(m)
	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, err = nc.ChanSubscribe("foo", ch)
	msg := <-ch
	fmt.Println(msg)
	// Unsubscribe
	sub.Unsubscribe()

	// Drain
	sub.Drain()

	// Requests
	msg, err = nc.Request("help", []byte("help me"), 10*time.Millisecond)

	// Replies
	nc.Subscribe("help", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I can help!"))
	})


	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	// Close connection
	nc.Close()
}
