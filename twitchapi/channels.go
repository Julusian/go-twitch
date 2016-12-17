package twitchapi

import (
	"fmt"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"

	"github.com/google/go-querystring/query"
)

// ChannelsMethod wraps up channel api calls
// https://dev.twitch.tv/docs/v5/reference/channels/
type ChannelsMethod struct {
	client *Client
}

// Channel queries the api for the Channel object for the specified channel
// If `id` is 0, returns the channel of authenticated user.
// https://dev.twitch.tv/docs/v5/reference/channels/#get-channel
// https://dev.twitch.tv/docs/v5/reference/channels/#get-channel-by-id
func (c *ChannelsMethod) Channel(id uint) (*twitch.Channel, error) {
	rel := "channel" // get authenticated channel
	if id > 0 {
		rel = fmt.Sprintf("channels/%d", id)
	}

	channel := new(twitch.Channel)
	_, err := c.client.Get(rel, channel)
	return channel, err
}

// TODO https://dev.twitch.tv/docs/v5/reference/channels/#update-channel

// Editors gets a list of users who are editors of the specified channel
// https://dev.twitch.tv/docs/v5/reference/channels/#get-channel-editors
func (c *ChannelsMethod) Editors(id uint) (*twitch.ChannelEditorsList, error) {
	rel := fmt.Sprintf("channels/%d/editors", id)

	editors := new(twitch.ChannelEditorsList)
	_, err := c.client.Get(rel, editors)
	return editors, err
}

// Follows returns a list of users who follow the specified channel
// https://dev.twitch.tv/docs/v5/reference/channels/#get-channel-followers
func (c *ChannelsMethod) Follows(id uint, opt *twitch.ListOptions) (*twitch.ChannelFollowsList, error) {
	rel := fmt.Sprintf("channels/%d/follows", id)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	follow := new(twitch.ChannelFollowsList)
	_, err := c.client.Get(rel, follow)
	return follow, err
}

// TODO https://dev.twitch.tv/docs/v5/reference/channels/#get-channel-teams

// Subscriptions returns a list of users who subscribe to the specified channel
// https://dev.twitch.tv/docs/v5/reference/channels/#get-channel-subscribers
func (c *ChannelsMethod) Subscriptions(id uint, opt *twitch.ListOptions) (*twitch.ChannelSubscriptionsList, error) {
	rel := fmt.Sprintf("channels/%d/subscriptions", id)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	subs := new(twitch.ChannelSubscriptionsList)
	_, err := c.client.Get(rel, subs)
	return subs, err
}

// TODO - make public
// Subscription checks if a user is subscribed to the specified channel
// https://dev.twitch.tv/docs/v5/reference/channels/#check-channel-subscription-by-user
func (c *ChannelsMethod) subscription(id uint, user uint) (*twitch.ChannelSubscription, error) {
	rel := fmt.Sprintf("channels/%d/subscriptions/%d", id, user)

	sub := new(twitch.ChannelSubscription)
	_, err := c.client.Get(rel, sub)
	return sub, err
}

// Videos returns a list of videos ordered by time of creation for channel `id`
// https://dev.twitch.tv/docs/v5/reference/channels/#get-channel-videos
func (c *ChannelsMethod) Videos(id uint, opt *twitch.ListOptions) (*twitch.VideosList, error) {
	rel := fmt.Sprintf("channels/%d/videos", id)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	videos := new(twitch.VideosList)
	_, err := c.client.Get(rel, videos)
	return videos, err
}

// TODO https://dev.twitch.tv/docs/v5/reference/channels/#start-channel-commercial

// TODO https://dev.twitch.tv/docs/v5/reference/channels/#reset-channel-stream-key
