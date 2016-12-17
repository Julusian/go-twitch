package twitch

type TopGamesList struct {
	Total int       `json:"_total,omitempty"`
	Top   []TopGame `json:"top,omitempty"`
}

type TopGame struct {
	Channels int  `json:"channels,omitempty"`
	Viewers  int  `json:"viewers,omitempty"`
	Game     Game `json:"game,omitempty"`
}

type Game struct {
	ID          int          `json:"_id,omitempty"`
	Box         PreviewImage `json:"box,omitempty"`
	GiantbombID int          `json:"giantbomb_id,omitempty"`
	Logo        PreviewImage `json:"logo,omitempty"`
	Name        string       `json:"name,omitempty"`
	Popularity  int          `json:"popularity,omitempty"`
}
