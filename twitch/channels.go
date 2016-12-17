// Channels methods of the twitch api.
// https://github.com/justintv/Twitch-API/blob/master/v3_resources/channels.md

package twitch

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

// used with GET /channels/:channel/videos
type VideosS struct {
	Total  int      `json:"_total,omitempty"`
	Videos []VideoS `json:"videos,omitempty"`
}

// used with GET /channels/:channel/editors
type EditorsS struct {
	Users []UserS `json:"users,omitempty"`
}

// used with GET /channels/:channel/follows
type FollowsS struct {
	Cursor  string    `json:"_cursor,omitempty"`
	Total   int       `json:"_total,omitempty"`
	Follows []FollowS `json:"follows,omitempty"`
}

type FollowS struct {
	CreatedAt     time.Time `json:"created_at,omitempty"`
	Notifications bool      `json:"notifications,omitempty"`
	User          UserS     `json:"user,omitempty"`
}

type SubsS struct {
	Total         int    `json:"_total,omitempty"`
	Subscriptions []SubS `json:"subscriptions,omitempty"`
}

type SubS struct {
	ID        string    `json:"_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	User      UserS     `json:"user,omitempty"`
}

type ChannelsMethod struct {
	client *Client
}

// Returns a channel object. If `name` is an empty string, returns the channel
// object of authenticated user.
func (c *ChannelsMethod) Channel(id uint) (*ChannelS, error) {
	rel := "channels" // get authenticated channel
	if id > 0 {
		rel = fmt.Sprintf("channels/%d", id)
	}

	channel := new(ChannelS)
	_, err := c.client.Get(rel, channel)
	return channel, err
}

// Returns a list of users who are editors of channel `name`.
func (c *ChannelsMethod) editors(name string) (*EditorsS, error) {
	rel := "channels/" + name + "/editors"

	editors := new(EditorsS)
	_, err := c.client.Get(rel, editors)
	return editors, err
}

// Returns a list of videos ordered by time of creation, starting with the most
// recent from channel `name`.
func (c *ChannelsMethod) Videos(id uint, opt *ListOptions) (*VideosS, error) {
	rel := fmt.Sprintf("channels/%d/videos", id)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	videos := new(VideosS)
	_, err := c.client.Get(rel, videos)
	return videos, err
}

// Returns a list of users the channel `name` is following.
func (c *ChannelsMethod) Follows(id uint, opt *ListOptions) (*FollowsS, error) {
	rel := fmt.Sprintf("channels/%d/follows", id)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	follow := new(FollowsS)
	_, err := c.client.Get(rel, follow)
	return follow, err
}

func (c *ChannelsMethod) subscriptions(id uint, opt *ListOptions) (*SubsS, error) {
	rel := fmt.Sprintf("channels/%d/subscriptions", id)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	subs := new(SubsS)
	_, err := c.client.Get(rel, subs)
	return subs, err
}

func (c *ChannelsMethod) subscription(id uint, user uint) (*SubS, error) {
	rel := fmt.Sprintf("channels/%d/subscriptions/%d", id, user)

	sub := new(SubS)
	_, err := c.client.Get(rel, sub)
	return sub, err
}

// TODO PUT /channels/:channel/

// TODO POST /channels/:channel/commercial
