package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	var servers string

	flag.StringVar(&servers, "s", "", "server names to allow for matchmaking")
	flag.Parse()

	serversArray := strings.Split(servers, ",")
	var blockedIps []string

	pops, err := Decode()
	if err != nil {
		fmt.Println(err)
	}

	for key, location := range pops {
		for k := range serversArray {
			if key != serversArray[k] {
				for _, relay := range location.Relay {
					blockedIps = append(blockedIps, relay.IP)
				}
			}
		}
	}

	if err := addRule(blockedIps); err != nil {
		fmt.Printf("Could not add firewall rule: %v", err)
		os.Exit(1)
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	done := make(chan bool, 1)
	go func() {
		<-exit
		done <- true
	}()

	fmt.Println("\nPress CTRL+C to exit...")
	<-done

	if err := removeRule(); err != nil {
		fmt.Printf("Could not remove firewall rule: %v", err)
	}

	os.Exit(0)
}
