package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var timeout time.Duration

func init() {
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "timeout")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	host := args[0]
	port := args[1]
	address := net.JoinHostPort(host, port)

	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		log.Fatalf("Connection error: %s\n", err)
	}

	defer func(client TelnetHandler) {
		err := client.Close()
		if err != nil {
			log.Fatalf("Close error: %s\n", err)
		}
	}(client)

	fmt.Fprintf(os.Stderr, "...Connected to %s\n", address)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go startSending(wg, client)
	go startReceiving(wg, client)
	wg.Wait()
}

func startReceiving(wg *sync.WaitGroup, client TelnetHandler) {
	func() {
		defer wg.Done()
		err := client.Receive()
		if err != nil {
			log.Fatalf("Receive error: %s\n", err)
		}
		fmt.Fprintf(os.Stderr, "...Connection was closed by peer\n")
	}()
}

func startSending(wg *sync.WaitGroup, client TelnetHandler) {
	func() {
		defer wg.Done()
		err := client.Send()
		if err != nil {
			log.Fatalf("Send error: %s\n", err)
		}
		fmt.Fprintf(os.Stderr, "...EOF\n")
	}()
}
