package twitch

import "time"

// Channel object
type ChannelS struct {
	ID                           int       `json:"_id,omitempty"`
	BroadcasterLanguage          string    `json:"broadcaster_language,omitempty"`
	CreatedAt                    time.Time `json:"created_at,omitempty"`
	DisplayName                  string    `json:"display_name,omitempty"`
	Email                        string    `json:"email,omitempty"` // Authenticated
	Followers                    int       `json:"followers,omitempty"`
	Game                         string    `json:"game,omitempty"`
	Language                     string    `json:"language,omitempty"`
	Logo                         string    `json:"logo,omitempty"`
	Mature                       bool      `json:"mature,omitempty"`
	Name                         string    `json:"name,omitempty"`
	Partner                      bool      `json:"partner,omitempty"`
	ProfileBanner                string    `json:"profile_banner,omitempty"`
	ProfileBannerBackgroundColor string    `json:"profile_banner_background_color,omitempty"`
	Status                       string    `json:"status,omitempty"`
	StreamKey                    string    `json:"stream_key,omitempty"` // Authenticated
	UpdatedAt                    time.Time `json:"updated_at,omitempty"`
	URL                          string    `json:"url,omitempty"`
	VideoBanner                  string    `json:"video_banner,omitempty"`
	Views                        int       `json:"views,omitempty"`
}

// Team object
type TeamS struct {
	ID          int       `json:"_id,omitempty"`
	Background  string    `json:"background,omitempty"`
	Banner      string    `json:"banner,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	Info        string    `json:"info,omitempty"`
	Logo        string    `json:"logo,omitempty"`
	Name        string    `json:"name,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Users       []UserS   `json:"users,omitempty"`
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

// User object
type UserS struct {
	ID               int       `json:"_id,omitempty"`
	Bio              string    `json:"bio,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	DisplayName      string    `json:"display_name,omitempty"`
	Email            string    `json:"email,omitempty"`          // Authenticated
	EmailVerified    bool      `json:"email_verified,omitempty"` // Authenticated
	Logo             string    `json:"logo,omitempty"`
	Name             string    `json:"name,omitempty"`
	Partnered        bool      `json:"partnered,omitempty"`         // Authenticated
	TwitterConnected bool      `json:"twitter_connected,omitempty"` // Authenticated
	Type             string    `json:"type,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`

	// Notifications UserNotificataions `json:"notifications,omitempty"` // Authenticated
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

type FStreamS struct {
	Image     string  `json:"image,omitempty"`
	Priority  int     `json:"priority,omitempty"`
	Scheduled bool    `json:"scheduled,omitempty"`
	Sponsored bool    `json:"sponsored,omitempty"`
	Stream    StreamS `json:"stream,omitempty"`
	Text      string  `json:"text,omitempty"`
	Title     string  `json:"title,omitempty"`
}

type PreviewS struct {
	Large    string `json:"large,omitempty"`
	Medium   string `json:"medium,omitempty"`
	Small    string `json:"small,omitempty"`
	Template string `json:"template,omitempty"`
}

type ListOptions struct {
	Game       string `url:"game,omitempty"`
	Channel    string `url:"channel,omitempty"`
	Direction  string `url:"direction,omitempty"`
	Period     string `url:"period,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty"`
	Cursor     string `url:"cursor,omitempty"`
	Embeddable *bool  `url:"embeddable,omitempty"`
	Hls        bool   `url:"hls,omitempty"`
	Live       *bool  `url:"live,omitempty"`
	Cache      int64  `url:"_,omitempty"`
	Broadcasts bool   `url:"broadcasts,omitempty"`
}
