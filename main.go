package main

import (
	"flag"
	"log"

	"github.com/inflame-ue/dlive/internal/sse"
)

var subscribed_entities = []string{
	"player_controller",
	"player_pawn",
	"team",
	"mid_boss",
}

func main() {
	matchID := flag.Int("match-id", 0, "Specifies the match ID")
	port := flag.Int("port", 3000, "Specify the port of the Deadlock API")
	flag.Parse()

	client := sse.NewMatchClient(*port, *matchID)
	log.Printf("establishing connection to API on port %d for match: %d", *port, *matchID)
	go client.EstablishConnection(true, subscribed_entities)

	for range client.Events {

	}
}
