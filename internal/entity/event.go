package entity

type Event interface {
	eventMarker()
}

type EventType string

const (
	Created EventType = "entity_create"
	Updated EventType = "entity_update"
	Deleted EventType = "entity_delete"
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
	Tick     int     `json:"tick"`
	GameTime float64 `json:"game_time"`
}

func (TickEndEvent) eventMarker() {}

type ChatMessageEvent struct {
	Tick     int     `json:"tick"`
	GameTime float64 `json:"game_time"`
	ChatMessage
}

func (ChatMessageEvent) eventMarker() {}
