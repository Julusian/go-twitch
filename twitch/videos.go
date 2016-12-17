package twitch

type VVideosS struct {
	Total  int      `json:"_total,omitempty"`
	Videos []VideoS `json:"vods,omitempty"`
}
