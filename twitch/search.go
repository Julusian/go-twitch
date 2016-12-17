package twitch

type GamesSearch struct {
	Games []Game `json:"games,omitempty"`
}

type ChannelsSearch struct {
	Total    int       `json:"_total,omitempty"`
	Channels []Channel `json:"channels,omitempty"`
}
