package main

import (
	"fmt"
	"log"
	"my_multicast/multicast"
	"os"
	"time"

	"github.com/urfave/cli"
)

const (
	defaultMulticastAddress = "239.0.0.0:9999"
)

type Machine struct {
	name string
	Age  int
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix)
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		address := c.Args().Get(0)
		if address == "" {
			address = defaultMulticastAddress
		}
		fmt.Printf("Broadcasting to %s\n", address)
		ping(address)
		return nil
	}
	app.Run(os.Args)
}

func ping(addr string) {
	conn, err := multicast.NewBroadcaster(addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn.Write([]byte("hello world\n"))
		time.Sleep(3 * time.Second)
	}
}
