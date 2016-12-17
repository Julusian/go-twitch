package twitchapi

import (
	"fmt"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"

	"github.com/google/go-querystring/query"
)

type UsersMethod struct {
	client *Client
}

func (u *UsersMethod) Name(name string) (*twitch.UserS, error) {
	rel := fmt.Sprintf("users?login=%s", name)

	usr := new(twitch.UserS)
	_, err := u.client.Get(rel, usr)
	return usr, err
}

// User returns a user object.
func (u *UsersMethod) User(id uint) (*twitch.UserS, error) {
	rel := "user" // get authenticated user
	if id > 0 {
		rel = fmt.Sprintf("users/%d", id)
	}

	usr := new(twitch.UserS)
	_, err := u.client.Get(rel, usr)
	return usr, err
}

func (u *UsersMethod) Blocks(id uint, opt *twitch.ListOptions) (*twitch.BlocksS, error) {
	rel := fmt.Sprintf("users/%d/blocks", id)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	blocks := new(twitch.BlocksS)
	_, err := u.client.Get(rel, blocks)
	return blocks, err
}

// Get a user's list of followed channels
func (u *UsersMethod) Follows(id uint, opt *twitch.ListOptions) (*twitch.UFollowsS, error) {
	rel := fmt.Sprintf("users/%d/follows/channels", id)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	follows := new(twitch.UFollowsS)
	_, err := u.client.Get(rel, follows)
	return follows, err
}

// Get status of follow relationship between user and target channel
func (u *UsersMethod) Follow(user uint, target uint) (*twitch.UFollowS, error) {
	rel := fmt.Sprintf("users/%d/follows/channels/%d", user, target)

	follow := new(twitch.UFollowS)
	_, err := u.client.Get(rel, follow)
	return follow, err
}

// TODO - not subscribed not handled
func (u *UsersMethod) subscription(user uint, channel uint) (*twitch.USubS, error) {
	rel := fmt.Sprintf("users/%d/subscriptions/%d", user, channel)

	follow := new(twitch.USubS)
	_, err := u.client.Get(rel, follow)
	return follow, err
}
