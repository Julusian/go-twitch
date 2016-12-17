package twitchapi

import (
	"fmt"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"

	"github.com/google/go-querystring/query"
)

type ChannelsMethod struct {
	client *Client
}

// Returns a channel object. If `name` is an empty string, returns the channel
// object of authenticated user.
func (c *ChannelsMethod) Channel(id uint) (*twitch.ChannelS, error) {
	rel := "channels" // get authenticated channel
	if id > 0 {
		rel = fmt.Sprintf("channels/%d", id)
	}

	channel := new(twitch.ChannelS)
	_, err := c.client.Get(rel, channel)
	return channel, err
}

// Returns a list of users who are editors of channel `name`.
func (c *ChannelsMethod) editors(name string) (*twitch.EditorsS, error) {
	rel := "channels/" + name + "/editors"

	editors := new(twitch.EditorsS)
	_, err := c.client.Get(rel, editors)
	return editors, err
}

// Returns a list of videos ordered by time of creation, starting with the most
// recent from channel `name`.
func (c *ChannelsMethod) Videos(id uint, opt *twitch.ListOptions) (*twitch.VideosS, error) {
	rel := fmt.Sprintf("channels/%d/videos", id)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	videos := new(twitch.VideosS)
	_, err := c.client.Get(rel, videos)
	return videos, err
}

// Returns a list of users the channel `name` is following.
func (c *ChannelsMethod) Follows(id uint, opt *twitch.ListOptions) (*twitch.FollowsS, error) {
	rel := fmt.Sprintf("channels/%d/follows", id)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	follow := new(twitch.FollowsS)
	_, err := c.client.Get(rel, follow)
	return follow, err
}

func (c *ChannelsMethod) subscriptions(id uint, opt *twitch.ListOptions) (*twitch.SubsS, error) {
	rel := fmt.Sprintf("channels/%d/subscriptions", id)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	subs := new(twitch.SubsS)
	_, err := c.client.Get(rel, subs)
	return subs, err
}

func (c *ChannelsMethod) subscription(id uint, user uint) (*twitch.SubS, error) {
	rel := fmt.Sprintf("channels/%d/subscriptions/%d", id, user)

	sub := new(twitch.SubS)
	_, err := c.client.Get(rel, sub)
	return sub, err
}

// TODO PUT /channels/:channel/

// TODO POST /channels/:channel/commercial
