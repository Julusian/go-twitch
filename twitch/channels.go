// Channels methods of the twitch api.
// https://github.com/justintv/Twitch-API/blob/master/v3_resources/channels.md

package twitch

import "time"

// used with GET /channels/:channel/videos
type VideosS struct {
	Total  int      `json:"_total,omitempty"`
	Videos []VideoS `json:"videos,omitempty"`
}

// used with GET /channels/:channel/editors
type EditorsS struct {
	Users []UserS `json:"users,omitempty"`
}

// used with GET /channels/:channel/follows
type FollowsS struct {
	Cursor  string    `json:"_cursor,omitempty"`
	Total   int       `json:"_total,omitempty"`
	Follows []FollowS `json:"follows,omitempty"`
}

type FollowS struct {
	CreatedAt     time.Time `json:"created_at,omitempty"`
	Notifications bool      `json:"notifications,omitempty"`
	User          UserS     `json:"user,omitempty"`
}

type SubsS struct {
	Total         int    `json:"_total,omitempty"`
	Subscriptions []SubS `json:"subscriptions,omitempty"`
}

type SubS struct {
	ID        string    `json:"_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	User      UserS     `json:"user,omitempty"`
}
