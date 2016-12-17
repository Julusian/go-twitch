package twitch

import "time"

// VideosList is a structure representing a list of videos
// https://dev.twitch.tv/docs/v5/reference/channels/#get-channel-videos
// https://dev.twitch.tv/docs/v5/reference/videos/#get-followed-videos
type VideosList struct {
	Total  int     `json:"_total,omitempty"`
	Videos []Video `json:"videos,omitempty"`
}

// TopVideosList is a list of the top vods
// https://dev.twitch.tv/docs/v5/reference/videos/#get-top-videos
type TopVideosList struct {
	Total  int     `json:"_total,omitempty"`
	Videos []Video `json:"vods,omitempty"`
}

// Video is a structure representing a twitch video or vod
// https://dev.twitch.tv/docs/v5/reference/videos/
// https://dev.twitch.tv/docs/v5/reference/videos/#get-video
type Video struct {
	ID              string             `json:"_id,omitempty"`
	BroadcastID     int                `json:"broadcast_id"`
	BroadcastType   string             `json:"broadcast_type,omitempty"`
	Channel         Channel            `json:"channel,omitempty"`
	CreatedAt       time.Time          `json:"created_at,omitempty"`
	Description     string             `json:"description,omitempty"`
	DescriptionHTML string             `json:"description_html,omitempty"`
	FPS             map[string]float64 `json:"fps,omitempty"`
	Game            string             `json:"game,omitempty"`
	Language        string             `json:"language,omitempty"`
	Length          int                `json:"length,omitempty"`
	Preview         PreviewImage       `json:"preview,omitempty"`
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
