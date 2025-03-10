package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Set up a channel to listen for interrupt or termination signals.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Load nftables ruleset from file "nft.ruleset"
	if err := LoadRuleset("nft.ruleset"); err != nil {
		log.Fatalf("Error loading nftables ruleset: %v", err)
	}
	log.Println("Successfully loaded nftables ruleset")

	// Wait until a signal is received.
	<-sigChan
	log.Println("Termination signal received. Flushing nftables ruleset...")

	// Flush the nftables ruleset when the program is stopping.
	if err := FlushRuleset(); err != nil {
		log.Printf("Error flushing nftables ruleset: %v", err)
	} else {
		log.Println("Successfully flushed nftables ruleset")
	}

	log.Println("Exiting program")
}
