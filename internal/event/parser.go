package event

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/inflame-ue/dlive/internal/sse"
)

type EventParser struct {
	RawEvents    chan *sse.RawEvent
	ParsedEvents chan Event
}

func NewEventParser(rawEvents chan *sse.RawEvent) *EventParser {
	return &EventParser{
		RawEvents:    rawEvents,
		ParsedEvents: make(chan Event),
	}
}

func parseEvent[T Event](data []byte) (T, error) {
	var event T
	if err := json.Unmarshal(data, &event); err != nil {
		return event, fmt.Errorf("unmarhsalling event: %w", err)
	}
	return event, nil
}

func sendParsedEvent[T Event](data []byte, ep *EventParser) {
	event, err := parseEvent[T](data)
	if err != nil {
		log.Print(err)
		return
	}
	ep.ParsedEvents <- event
}

func (ep *EventParser) ParseRawEvents() {
	for event := range ep.RawEvents {
		log.Print("parsing raw events")
		rawEventType := strings.TrimSpace(event.EventType)

		switch rawEventType {
		case "end":
			sendParsedEvent[GameEndEvent](event.Data, ep)
			continue
		case "tick_end":
			sendParsedEvent[TickEndEvent](event.Data, ep)
			continue
		case "chat_message":
			sendParsedEvent[ChatMessageEvent](event.Data, ep)
			continue
		}

		entity, _, _ := strings.Cut(rawEventType, "_entity_")
		switch entity {
		case "player_controller":
			sendParsedEvent[PlayerControllerEvent](event.Data, ep)
		case "player_pawn":
			sendParsedEvent[PlayerPawnEvent](event.Data, ep)
		case "team":
			sendParsedEvent[TeamEvent](event.Data, ep)
		case "mid_boss":
			sendParsedEvent[MidBossEvent](event.Data, ep)
		default:
			log.Print("unknown entity type")
			continue
		}
	}
}
