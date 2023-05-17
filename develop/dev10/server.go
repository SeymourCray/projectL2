package main

import (
	"fmt"
	"github.com/aprice/telnet"
	"log"
)

func main() {
	svr := telnet.NewServer(":9999", telnet.HandleFunc(func(c *telnet.Connection) {
		log.Printf("Connection received: %s", c.RemoteAddr())

		for {
			s := make([]byte, 100)
			if _, err := c.Read(s); err == nil {
				c.Write(s)
				fmt.Print(string(s))
			}
		}
	}))

	svr.ListenAndServe()
}
