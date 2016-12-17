package twitch

type TopsS struct {
	Total int    `json:"_total,omitempty"`
	Top   []TopS `json:"top,omitempty"`
}

type TopS struct {
	Channels int   `json:"channels,omitempty"`
	Viewers  int   `json:"viewers,omitempty"`
	Game     GameS `json:"game,omitempty"`
}

type GameS struct {
	ID          int      `json:"_id,omitempty"`
	Box         PreviewS `json:"box,omitempty"`
	GiantbombID int      `json:"giantbomb_id,omitempty"`
	Logo        PreviewS `json:"logo,omitempty"`
	Name        string   `json:"name,omitempty"`
	Popularity  int      `json:"popularity,omitempty"`
}
