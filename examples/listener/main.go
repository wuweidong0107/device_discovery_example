package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"my_multicast/multicast"
	"net"
	"os"

	"github.com/urfave/cli"
)

const (
	defaultMulticastAddress = "239.0.0.0:9999"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix)
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		address := c.Args().Get(0)
		if address == "" {
			address = defaultMulticastAddress
		}
		fmt.Printf("Listening to %s\n", address)
		multicast.Listen(address, msgHandler)
		return nil
	}
	app.Run(os.Args)
}

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))
}
