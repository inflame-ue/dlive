package event

type Event interface {
	eventMarker()
}

type EventType string

const (
	Unknown EventType = "unknown"
	Created EventType = "entity_create"
	Updated EventType = "entity_update"
	Deleted EventType = "entity_delete"
	TickEnd EventType = "tick_end"
	Message EventType = "chat_message"
	GameEnd EventType = "end"
)

type PlayerControllerEvent struct {
	Event       EventType `json:"event_type"`
	Tick        int       `json:"tick"`
	GameTime    float64   `json:"game_time"`
	EntityIndex int       `json:"entity_index"`
	PlayerController
}

func (PlayerControllerEvent) eventMarker() {}

type PlayerPawnEvent struct {
	Event       EventType `json:"event_type"`
	Tick        int       `json:"tick"`
	GameTime    float64   `json:"game_time"`
	EntityIndex int       `json:"entity_index"`
	PlayerPawn
}

func (PlayerPawnEvent) eventMarker() {}

type TeamEvent struct {
	Event       EventType `json:"event_type"`
	Tick        int       `json:"tick"`
	GameTime    float64   `json:"game_time"`
	EntityIndex int       `json:"entity_index"`
	Team
}

func (TeamEvent) eventMarker() {}

type MidBossEvent struct {
	Event       EventType `json:"event_type"`
	Tick        int       `json:"tick"`
	GameTime    float64   `json:"game_time"`
	EntityIndex int       `json:"entity_index"`
	MidBoss
}

func (MidBossEvent) eventMarker() {}

type TickEndEvent struct {
	Event    EventType `json:"event_type"`
	Tick     int       `json:"tick"`
	GameTime float64   `json:"game_time"`
}

func (TickEndEvent) eventMarker() {}

type ChatMessageEvent struct {
	Event    EventType `json:"event_type"`
	Tick     int       `json:"tick"`
	GameTime float64   `json:"game_time"`
	ChatMessage
}

func (ChatMessageEvent) eventMarker() {}

type GameEndEvent struct {
	Event EventType `json:"event_type"`
}

func (GameEndEvent) eventMarker() {}
