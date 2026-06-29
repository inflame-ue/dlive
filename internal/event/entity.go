package event

type PlayerController struct {
	SteamID              int64   `json:"steam_id"`
	SteamName            string  `json:"steam_name"`
	Team                 int     `json:"team"`
	HeroID               int     `json:"hero_id"`
	HeroBadgeXp          int     `json:"hero_badge_xp"`
	PlayerSlot           int     `json:"player_slot"`
	Rank                 int     `json:"rank"`
	AssignedLane         int     `json:"assigned_lane"`
	OriginalAssignedLane int     `json:"original_assigned_lane"`
	NetWorth             int     `json:"net_worth"`
	UltimateTrained      bool    `json:"ultimate_trained"`
	Kills                int     `json:"kills"`
	Assists              int     `json:"assists"`
	Deaths               int     `json:"deaths"`
	Denies               int     `json:"denies"`
	LastHits             int     `json:"last_hits"`
	HeroHealing          int     `json:"hero_healing"`
	SelfHealing          int     `json:"self_healing"`
	HeroDamage           int     `json:"hero_damage"`
	ObjectiveDamage      int     `json:"objective_damage"`
	UltimateCooldownEnd  float64 `json:"ultimate_cooldown_end"`
	Upgrades             []int   `json:"upgrades"`
}

type PlayerPawn struct {
	Controller             int       `json:"controller"`
	Team                   int      `json:"team"`
	HeroID                 *int       `json:"hero_id"`
	HeroBuildID            int       `json:"hero_build_id"`
	HeroBuildSerialized    []int     `json:"hero_build_serialized"`
	QuickbuyQueue          []int     `json:"quickbuy_queue"`
	QuickbuyAutoPurchase   bool      `json:"quickbuy_auto_purchase"`
	QuickbuyAutoQueueBuild bool      `json:"quickbuy_auto_queue_build"`
	Level                  int       `json:"level"`
	MaxHealth              int       `json:"max_health"`
	Health                 int       `json:"health"`
	Position               []float64 `json:"position"`
}

type Team struct {
	Team         int    `json:"team"`
	Score        int    `json:"score"`
	Teamname     string `json:"teamname"`
	FlexUnlocked int    `json:"flex_unlocked"`
}

type MidBoss struct {
	Health     int       `json:"health"`
	MaxHealth  int       `json:"max_health"`
	CreateTime float64   `json:"create_time"`
	Team       int       `json:"team"`
	Position   []float64 `json:"position"`
}

type ChatMessage struct {
	SteamName string `json:"steam_name"`
	SteamID   int64  `json:"steam_id"`
	Text      string `json:"text"`
	AllChat   bool   `json:"all_chat"`
}
