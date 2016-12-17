package twitch

import "time"

type BlocksS struct {
	Total  int      `json:"_total,omitempty"`
	Blocks []BlockS `json:"blocks,omitempty"`
}

type BlockS struct {
	ID        int       `json:"_id,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	User      UserS     `json:"user,omitempty"`
}

type UFollowsS struct {
	Total   int        `json:"_total,omitempty"`
	Follows []UFollowS `json:"follows,omitempty"`
}

type UFollowS struct {
	CreatedAt     string   `json:"created_at,omitempty"`
	Notifications bool     `json:"notifications,omitempty"`
	Channel       ChannelS `json:"channel,omitempty"`
}

type USubS struct {
	ID        string    `json:"_id,omitempty"`
	Channel   ChannelS  `json:"channel,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
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
