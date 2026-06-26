package entity

import "github.com/inflame-ue/dlive/internal/sse"

type Parser struct {
	RawEvents    chan sse.RawEvent
	ParsedEvents chan Event
}
