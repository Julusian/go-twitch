package twitchapi

import (
	"fmt"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"

	"github.com/google/go-querystring/query"
)

type VideosMethod struct {
	client *Client
}

func (v *VideosMethod) Id(id uint) (*twitch.VideoS, error) {
	rel := fmt.Sprintf("videos/%d", id)

	video := new(twitch.VideoS)
	_, err := v.client.Get(rel, video)
	return video, err
}

func (v *VideosMethod) Top(opt *twitch.ListOptions) (*twitch.VVideosS, error) {
	rel := "videos/top"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	videos := new(twitch.VVideosS)
	_, err := v.client.Get(rel, videos)
	return videos, err
}

func (v *VideosMethod) Followed(opt *twitch.ListOptions) (*twitch.VideosS, error) {
	rel := "videos/followed"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	videos := new(twitch.VideosS)
	_, err := v.client.Get(rel, videos)
	return videos, err
}