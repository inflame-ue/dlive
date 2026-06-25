package sse

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type RawEvent struct {
	EventType string
	Data      []byte
}

type MatchClient struct {
	port    int
	matchID int
	Events  chan *RawEvent
}

func NewMatchClient(deadlockAPIPort, deadlockMatchID int) *MatchClient {
	return &MatchClient{
		port:    deadlockAPIPort,
		matchID: deadlockMatchID,
		Events:  make(chan *RawEvent),
	}
}

func parseEvent(eventString []string) (*RawEvent, error) {
	if len(eventString) < 2 {
		return nil, fmt.Errorf("the server-sent event must have two parts: event type and data bytes")
	}

	eventLine, dataLine := eventString[0], eventString[1]
	_, eventType, _ := strings.Cut(eventLine, ":")
	eventType = strings.TrimSpace(eventType)
	
	_, data, _ := strings.Cut(dataLine, ":")
	data = strings.TrimSpace(data)

	return &RawEvent{
		EventType: eventType,
		Data:      []byte(data),
	}, nil
}

func (mc *MatchClient) EstablishConnection(subscribeChatMessages bool, subscribedEntities []string) {
	url := fmt.Sprintf("http://localhost:%d/v1/matches/%d/live/demo/events?subscribed_chat_messages=%s&subscribed_entities=%s",
		mc.port, mc.matchID, strconv.FormatBool(subscribeChatMessages), strings.Join(subscribedEntities, ","))

	log.Printf("establishing connection for url %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("establishing connection: %v", err)
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	eventString := make([]string, 0, 2)
	for scanner.Scan() {
		if scanner.Text() == "" {
			event, err := parseEvent(eventString)
			if err != nil {
				log.Print(err)
				continue
			}
			log.Printf("Event Type: %s, Event Data: %s", event.EventType, event.Data)
			mc.Events <- event
			eventString = nil
		} else {
			eventString = append(eventString, scanner.Text())
		}
	}

	close(mc.Events)
	if err := scanner.Err(); err != nil {
		log.Printf("scanner: %v", err)
		return
	}
}
