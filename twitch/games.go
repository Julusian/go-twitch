package twitch

// TopGamesList is a list of the top games being played on twitch
// https://dev.twitch.tv/docs/v5/reference/games/#get-top-games
type TopGamesList struct {
	Total int       `json:"_total,omitempty"`
	Top   []TopGame `json:"top,omitempty"`
}

// TopGame is an entry of TopGamesList
// https://dev.twitch.tv/docs/v5/reference/games/#get-top-games
type TopGame struct {
	Channels int  `json:"channels,omitempty"`
	Viewers  int  `json:"viewers,omitempty"`
	Game     Game `json:"game,omitempty"`
}

// Game is a structure representing a game on twitch
// https://dev.twitch.tv/docs/v5/reference/games/
type Game struct {
	ID          int          `json:"_id,omitempty"`
	Box         PreviewImage `json:"box,omitempty"`
	GiantbombID int          `json:"giantbomb_id,omitempty"`
	Logo        PreviewImage `json:"logo,omitempty"`
	Name        string       `json:"name,omitempty"`
	Popularity  int          `json:"popularity,omitempty"`
}
