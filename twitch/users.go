package twitch

import "time"

type UserBlocksList struct {
	Total  int         `json:"_total,omitempty"`
	Blocks []UserBlock `json:"blocks,omitempty"`
}

type UserBlock struct {
	ID        int       `json:"_id,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	User      User      `json:"user,omitempty"`
}

type UserFollowsList struct {
	Total   int          `json:"_total,omitempty"`
	Follows []UserFollow `json:"follows,omitempty"`
}

type UserFollow struct {
	CreatedAt     string  `json:"created_at,omitempty"`
	Notifications bool    `json:"notifications,omitempty"`
	Channel       Channel `json:"channel,omitempty"`
}

type UserSubscription struct {
	ID        string    `json:"_id,omitempty"`
	Channel   Channel   `json:"channel,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// User object
type User struct {
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
