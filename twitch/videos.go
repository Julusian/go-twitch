package twitch

import "time"

type VVideosS struct {
	Total  int      `json:"_total,omitempty"`
	Videos []VideoS `json:"vods,omitempty"`
}

// Video object
type VideoS struct {
	ID              string             `json:"_id,omitempty"`
	BroadcastID     int                `json:"broadcast_id"`
	BroadcastType   string             `json:"broadcast_type,omitempty"`
	Channel         ChannelS           `json:"channel,omitempty"`
	CreatedAt       time.Time          `json:"created_at,omitempty"`
	Description     string             `json:"description,omitempty"`
	DescriptionHTML string             `json:"description_html,omitempty"`
	FPS             map[string]float64 `json:"fps,omitempty"`
	Game            string             `json:"game,omitempty"`
	Language        string             `json:"language,omitempty"`
	Length          int                `json:"length,omitempty"`
	Preview         PreviewS           `json:"preview,omitempty"`
	PublishedAt     time.Time          `json:"published_at,omitempty"`
	Resolutions     map[string]string  `json:"resolutions,omitempty"`
	Status          string             `json:"status,omitempty"`
	TagList         string             `json:"tag_list,omitempty"`
	Title           string             `json:"title,omitempty"`
	URL             string             `json:"url,omitempty"`
	Viewable        string             `json:"viewable,omitempty"`
	ViewableAt      string             `json:"viewable_at,omitempty"`
	Views           int                `json:"views,omitempty"`

	// TODO - define these types
	// Thumbnails  VideoThumbnails   `json:"thumbnails,omitempty"`
}
