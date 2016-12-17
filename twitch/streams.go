// Streams methods of the twitch api.
// https://github.com/justintv/Twitch-API/blob/master/v3_resources/streams.md

package twitch

// used with GET /streams/:channel/
type SChannelS struct {
	Stream *StreamS `json:"stream,omitempty"`
}

// used with GET /streams
type StreamsS struct {
	Total   int       `json:"_total,omitempty"`
	Streams []StreamS `json:"streams,omitempty"`
}

// used with GET /streams/featured
type FeaturedS struct {
	Featured []FStreamS `json:"featured,omitempty"`
}

// used with GET /streams/summary
type SummaryS struct {
	Channels int `json:"channels,omitempty"`
	Viewers  int `json:"viewers,omitempty"`
}

// used with GET /streams/followed
type FollowedS struct {
	Total   int       `json:"_total,omitempty"`
	Streams []StreamS `json:"streams,omitempty"`
}
