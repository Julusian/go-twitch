package twitch

type SGamesS struct {
	Games []GameS `json:"games,omitempty"`
}

type SChannelsS struct {
	Total    int        `json:"_total,omitempty"`
	Channels []ChannelS `json:"channels,omitempty"`
}
