package twitch

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

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

type UsersMethod struct {
	client *Client
}

func (u *UsersMethod) Name(name string) (*UserS, error) {
	rel := fmt.Sprintf("users?login=%s", name)

	usr := new(UserS)
	_, err := u.client.Get(rel, usr)
	return usr, err
}

// User returns a user object.
func (u *UsersMethod) User(id uint) (*UserS, error) {
	rel := "user" // get authenticated user
	if id > 0 {
		rel = fmt.Sprintf("users/%d", id)
	}

	usr := new(UserS)
	_, err := u.client.Get(rel, usr)
	return usr, err
}

func (u *UsersMethod) Blocks(id uint, opt *ListOptions) (*BlocksS, error) {
	rel := fmt.Sprintf("users/%d/blocks", id)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	blocks := new(BlocksS)
	_, err := u.client.Get(rel, blocks)
	return blocks, err
}

// Get a user's list of followed channels
func (u *UsersMethod) Follows(id uint, opt *ListOptions) (*UFollowsS, error) {
	rel := fmt.Sprintf("users/%d/follows/channels", id)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	follows := new(UFollowsS)
	_, err := u.client.Get(rel, follows)
	return follows, err
}

// Get status of follow relationship between user and target channel
func (u *UsersMethod) Follow(user uint, target uint) (*UFollowS, error) {
	rel := fmt.Sprintf("users/%d/follows/channels/%d", user, target)

	follow := new(UFollowS)
	_, err := u.client.Get(rel, follow)
	return follow, err
}

// TODO - not subscribed not handled
func (u *UsersMethod) subscription(user uint, channel uint) (*USubS, error) {
	rel := fmt.Sprintf("users/%d/subscriptions/%d", user, channel)

	follow := new(USubS)
	_, err := u.client.Get(rel, follow)
	return follow, err
}
