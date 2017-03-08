package twitch

import "time"

// ChannelStream is the result when querying the current stream for a channel
// https://dev.twitch.tv/docs/v5/reference/streams/#get-stream-by-channel
type ChannelStream struct {
	Stream *Stream `json:"stream,omitempty"`
}

// StreamsList is a list of streams
// https://dev.twitch.tv/docs/v5/reference/streams/#get-all-streams
// https://dev.twitch.tv/docs/v5/reference/streams/#get-followed-streams
type StreamsList struct {
	Total   int      `json:"_total,omitempty"`
	Streams []Stream `json:"streams,omitempty"`
}

// FeaturedStreamsList is a list of featured streams
// https://dev.twitch.tv/docs/v5/reference/streams/#get-featured-streams
type FeaturedStreamsList struct {
	Featured []FeaturedStream `json:"featured,omitempty"`
}

// FeaturedStream is an entry of FeaturedStreamsList
// https://dev.twitch.tv/docs/v5/reference/streams/#get-featured-streams
type FeaturedStream struct {
	Image     string `json:"image,omitempty"`
	Priority  int    `json:"priority,omitempty"`
	Scheduled bool   `json:"scheduled,omitempty"`
	Sponsored bool   `json:"sponsored,omitempty"`
	Stream    Stream `json:"stream,omitempty"`
	Text      string `json:"text,omitempty"`
	Title     string `json:"title,omitempty"`
}

// StreamSummary is a summary of all the live streams on twitch
// https://dev.twitch.tv/docs/v5/reference/streams/#get-streams-summary
type StreamSummary struct {
	Channels int `json:"channels,omitempty"`
	Viewers  int `json:"viewers,omitempty"`
}

// Stream is a structure representing a stream on twitch
// https://dev.twitch.tv/docs/v5/reference/streams/
type Stream struct {
	ID          TwitchID     `json:"_id,omitempty"`
	Game        string       `json:"game,omitempty"`
	Viewers     int          `json:"viewers,omitempty"`
	VideoHeight int          `json:"video_height,omitempty"`
	AverageFPS  float64      `json:"average_fps,omitempty"`
	Delay       int          `json:"delay,omitempty"`
	CreatedAt   time.Time    `json:"created_at,omitempty"`
	IsPlaylist  bool         `json:"is_playlist,omitempty"`
	Preview     PreviewImage `json:"preview,omitempty"`
	Channel     Channel      `json:"channel,omitempty"`

	// Name        string   `json:"name,omitempty"`
	// Broadcaster string   `json:"broadcaster,omitempty"`
	// Geo         string   `json:"geo,omitempty"`
	// Status      string   `json:"status,omitempty"`
}
