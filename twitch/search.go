package twitch

// GamesSearch is the result when searching for games
// https://dev.twitch.tv/docs/v5/reference/search/#search-games
type GamesSearch struct {
	Games []Game `json:"games,omitempty"`
}

// ChannelsSearch is the result when searching for channels
// https://dev.twitch.tv/docs/v5/reference/search/#search-channels
type ChannelsSearch struct {
	Total    int       `json:"_total,omitempty"`
	Channels []Channel `json:"channels,omitempty"`
}
