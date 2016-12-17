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
