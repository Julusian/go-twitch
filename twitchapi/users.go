package twitchapi

import (
	"fmt"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"

	"github.com/google/go-querystring/query"
)

type UsersMethod struct {
	client *Client
}

func (u *UsersMethod) Name(name string) (*twitch.User, error) {
	rel := fmt.Sprintf("users?login=%s", name)

	usr := new(twitch.User)
	_, err := u.client.Get(rel, usr)
	return usr, err
}

// User returns a user object.
func (u *UsersMethod) User(id uint) (*twitch.User, error) {
	rel := "user" // get authenticated user
	if id > 0 {
		rel = fmt.Sprintf("users/%d", id)
	}

	usr := new(twitch.User)
	_, err := u.client.Get(rel, usr)
	return usr, err
}

func (u *UsersMethod) Blocks(id uint, opt *twitch.ListOptions) (*twitch.UserBlocksList, error) {
	rel := fmt.Sprintf("users/%d/blocks", id)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	blocks := new(twitch.UserBlocksList)
	_, err := u.client.Get(rel, blocks)
	return blocks, err
}

// Get a user's list of followed channels
func (u *UsersMethod) Follows(id uint, opt *twitch.ListOptions) (*twitch.UserFollowsList, error) {
	rel := fmt.Sprintf("users/%d/follows/channels", id)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	follows := new(twitch.UserFollowsList)
	_, err := u.client.Get(rel, follows)
	return follows, err
}

// Get status of follow relationship between user and target channel
func (u *UsersMethod) Follow(user uint, target uint) (*twitch.UserFollow, error) {
	rel := fmt.Sprintf("users/%d/follows/channels/%d", user, target)

	follow := new(twitch.UserFollow)
	_, err := u.client.Get(rel, follow)
	return follow, err
}

// TODO - not subscribed not handled
func (u *UsersMethod) subscription(user uint, channel uint) (*twitch.UserSubscription, error) {
	rel := fmt.Sprintf("users/%d/subscriptions/%d", user, channel)

	follow := new(twitch.UserSubscription)
	_, err := u.client.Get(rel, follow)
	return follow, err
}
