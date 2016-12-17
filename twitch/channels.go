package twitch

import "time"

// ChannelEditorsList is a structure representing the editors of a channel
// https://dev.twitch.tv/docs/v5/reference/channels/#get-channel-editors
type ChannelEditorsList struct {
	Users []User `json:"users,omitempty"`
}

// ChannelFollowsList is a structure representing the followers of a channel
// https://dev.twitch.tv/docs/v5/reference/channels/#get-channel-followers
type ChannelFollowsList struct {
	Cursor  string          `json:"_cursor,omitempty"`
	Total   int             `json:"_total,omitempty"`
	Follows []ChannelFollow `json:"follows,omitempty"`
}

// ChannelFollow is an element of ChannelFollowsList
// https://dev.twitch.tv/docs/v5/reference/channels/#get-channel-followers
type ChannelFollow struct {
	CreatedAt     time.Time `json:"created_at,omitempty"`
	Notifications bool      `json:"notifications,omitempty"`
	User          User      `json:"user,omitempty"`
}

// ChannelSubscriptionsList is a structure representing the subscribers of a channel
// https://dev.twitch.tv/docs/v5/reference/channels/#get-channel-subscribers
type ChannelSubscriptionsList struct {
	Total         int                   `json:"_total,omitempty"`
	Subscriptions []ChannelSubscription `json:"subscriptions,omitempty"`
}

// ChannelSubscription is an element of ChannelSubscriptionsList
// https://dev.twitch.tv/docs/v5/reference/channels/#get-channel-subscribers
type ChannelSubscription struct {
	ID        string    `json:"_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	User      User      `json:"user,omitempty"`
}

// Channel is a structure representing a twitch channel
// https://dev.twitch.tv/docs/v5/reference/channels/#get-channel
// https://dev.twitch.tv/docs/v5/reference/channels/#get-channel-by-id
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
