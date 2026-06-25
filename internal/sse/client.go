package sse

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type MatchClient struct {
	port    int
	matchID int
	Events  chan []byte
}

func NewMatchClient(deadlockAPIPort, deadlockMatchID int) *MatchClient {
	return &MatchClient{
		port:    deadlockAPIPort,
		matchID: deadlockMatchID,
		Events:  make(chan []byte),
	}
}

func (mc *MatchClient) EstablishConnection(subscribeChatMessages bool, subscribedEntities []string) {
	url := fmt.Sprintf("http://localhost:%d/v1/matches/%d/live/demo/events?subscribed_chat_messages=%s&subscribed_entities=%s",
		mc.port, mc.matchID, strconv.FormatBool(subscribeChatMessages), strings.Join(subscribedEntities, ","))

	log.Printf("establishing connection for url %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("establishing connection: %w", err)
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		event := scanner.Bytes()
		log.Printf("scanned event: %s", event)
		mc.Events <- event
	}
}
