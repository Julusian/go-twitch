// Streams methods of the twitch api.
// https://github.com/justintv/Twitch-API/blob/master/v3_resources/streams.md

package twitch

import "time"

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

type FStreamS struct {
	Image     string  `json:"image,omitempty"`
	Priority  int     `json:"priority,omitempty"`
	Scheduled bool    `json:"scheduled,omitempty"`
	Sponsored bool    `json:"sponsored,omitempty"`
	Stream    StreamS `json:"stream,omitempty"`
	Text      string  `json:"text,omitempty"`
	Title     string  `json:"title,omitempty"`
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

// Stream oject
type StreamS struct {
	ID          int       `json:"_id,omitempty"`
	Game        string    `json:"game,omitempty"`
	Viewers     int       `json:"viewers,omitempty"`
	VideoHeight int       `json:"video_height,omitempty"`
	AverageFPS  int       `json:"average_fps,omitempty"`
	Delay       int       `json:"delay,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	IsPlaylist  bool      `json:"is_playlist,omitempty"`
	Preview     PreviewS  `json:"preview,omitempty"`
	Channel     ChannelS  `json:"channel,omitempty"`

	// Name        string   `json:"name,omitempty"`
	// Broadcaster string   `json:"broadcaster,omitempty"`
	// Geo         string   `json:"geo,omitempty"`
	// Status      string   `json:"status,omitempty"`
}
