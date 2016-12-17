package twitchapi

import (
	"fmt"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"

	"github.com/google/go-querystring/query"
)

// UsersMethod wraps up user api calls
// https://dev.twitch.tv/docs/v5/reference/users/
type UsersMethod struct {
	client *Client
}

// Name returns a user object for a user specified by their login name
func (u *UsersMethod) Name(name string) (*twitch.User, error) {
	rel := fmt.Sprintf("users?login=%s", name)

	usr := new(twitch.User)
	_, err := u.client.Get(rel, usr)
	return usr, err
}

// User returns a user object for the specified user or the authenticated user if id is 0
// https://dev.twitch.tv/docs/v5/reference/users/#get-user
// https://dev.twitch.tv/docs/v5/reference/users/#get-user-by-id
func (u *UsersMethod) User(id uint) (*twitch.User, error) {
	rel := "user" // get authenticated user
	if id > 0 {
		rel = fmt.Sprintf("users/%d", id)
	}

	usr := new(twitch.User)
	_, err := u.client.Get(rel, usr)
	return usr, err
}

// TODO https://dev.twitch.tv/docs/v5/reference/users/#get-user-emotes

// TODO - not subscribed not handled
// Subscription checks whether a user is subscribed to a channel
// https://dev.twitch.tv/docs/v5/reference/users/#check-user-subscription-by-channel
func (u *UsersMethod) subscription(user uint, channel uint) (*twitch.UserSubscription, error) {
	rel := fmt.Sprintf("users/%d/subscriptions/%d", user, channel)

	follow := new(twitch.UserSubscription)
	_, err := u.client.Get(rel, follow)
	return follow, err
}

// Follows returns a list of channels that the user is following
// https://dev.twitch.tv/docs/v5/reference/users/#get-user-follows
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

// Follow gets the status of follow relationship between user and target channel
// https://dev.twitch.tv/docs/v5/reference/users/#check-user-follows-by-channel
func (u *UsersMethod) Follow(user uint, target uint) (*twitch.UserFollow, error) {
	rel := fmt.Sprintf("users/%d/follows/channels/%d", user, target)

	follow := new(twitch.UserFollow)
	_, err := u.client.Get(rel, follow)
	return follow, err
}

// TODO https://dev.twitch.tv/docs/v5/reference/users/#follow-channel
// TODO https://dev.twitch.tv/docs/v5/reference/users/#unfollow-channel

// Blocks returns a list of users blocked by the specified user
// https://dev.twitch.tv/docs/v5/reference/users/#get-user-block-list
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

// TODO https://dev.twitch.tv/docs/v5/reference/users/#block-user
// TODO https://dev.twitch.tv/docs/v5/reference/users/#unblock-user
