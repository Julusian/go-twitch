// Channels methods of the twitch api.
// https://github.com/justintv/Twitch-API/blob/master/v3_resources/channels.md

package twitch

import "time"

// used with GET /channels/:channel/videos
type ChannelVideosList struct {
	Total  int     `json:"_total,omitempty"`
	Videos []Video `json:"videos,omitempty"`
}

// used with GET /channels/:channel/editors
type ChannelEditorsList struct {
	Users []User `json:"users,omitempty"`
}

// used with GET /channels/:channel/follows
type ChannelFollowsList struct {
	Cursor  string          `json:"_cursor,omitempty"`
	Total   int             `json:"_total,omitempty"`
	Follows []ChannelFollow `json:"follows,omitempty"`
}

type ChannelFollow struct {
	CreatedAt     time.Time `json:"created_at,omitempty"`
	Notifications bool      `json:"notifications,omitempty"`
	User          User      `json:"user,omitempty"`
}

type ChannelSubscriptionsList struct {
	Total         int                   `json:"_total,omitempty"`
	Subscriptions []ChannelSubscription `json:"subscriptions,omitempty"`
}

type ChannelSubscription struct {
	ID        string    `json:"_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	User      User      `json:"user,omitempty"`
}

// Channel object
type Channel struct {
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
