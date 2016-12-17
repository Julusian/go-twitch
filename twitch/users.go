package twitch

import "time"

// UserBlocksList is a list of users that the logged in user has blocked
// https://dev.twitch.tv/docs/v5/reference/users/#get-user-block-list
type UserBlocksList struct {
	Total  int         `json:"_total,omitempty"`
	Blocks []UserBlock `json:"blocks,omitempty"`
}

// UserBlock is an entry of UserBlocksList
// https://dev.twitch.tv/docs/v5/reference/users/#get-user-block-list
type UserBlock struct {
	ID        int       `json:"_id,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	User      User      `json:"user,omitempty"`
}

// UserFollowsList is a list of channels followed by a user
// https://dev.twitch.tv/docs/v5/reference/users/#get-user-follows
type UserFollowsList struct {
	Total   int          `json:"_total,omitempty"`
	Follows []UserFollow `json:"follows,omitempty"`
}

// UserFollow is an entry of UserFollowsList
// https://dev.twitch.tv/docs/v5/reference/users/#get-user-follows
type UserFollow struct {
	CreatedAt     string  `json:"created_at,omitempty"`
	Notifications bool    `json:"notifications,omitempty"`
	Channel       Channel `json:"channel,omitempty"`
}

// UserSubscription is a structure representing a subscription of a user to a channel
// https://dev.twitch.tv/docs/v5/reference/users/#check-user-subscription-by-channel
type UserSubscription struct {
	ID        string    `json:"_id,omitempty"`
	Channel   Channel   `json:"channel,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// User represents a user on twitch
// https://dev.twitch.tv/docs/v5/reference/users/
// https://dev.twitch.tv/docs/v5/reference/users/#get-user
// https://dev.twitch.tv/docs/v5/reference/users/#get-user-by-id
type User struct {
	ID               TwitchID  `json:"_id,omitempty"`
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
