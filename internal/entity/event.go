package entity

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

type PlayerPawnEvent struct {
	Event       EventType `json:"event_type"`
	Tick        int       `json:"tick"`
	GameTime    float64   `json:"game_time"`
	EntityIndex int       `json:"entity_index"`
	PlayerPawn
}

type TeamEvent struct {
	Event       EventType `json:"event_type"`
	Tick        int       `json:"tick"`
	GameTime    float64   `json:"game_time"`
	EntityIndex int       `json:"entity_index"`
	Team
}

type MidbossEvent struct {
	Event       EventType `json:"event_type"`
	Tick        int       `json:"tick"`
	GameTime    float64   `json:"game_time"`
	EntityIndex int       `json:"entity_index"`
	MidBoss
}

type TickEndEvent struct {
	Tick     int     `json:"tick"`
	GameTime float64 `json:"game_time"`
}

type ChatMessageEvent struct {
	Tick     int     `json:"tick"`
	GameTime float64 `json:"game_time"`
	ChatMessage
}
