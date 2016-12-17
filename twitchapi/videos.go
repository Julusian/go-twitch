package twitchapi

import (
	"fmt"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"

	"github.com/google/go-querystring/query"
)

// VideosMethod wraps up video api calls
// https://dev.twitch.tv/docs/v5/reference/videos/
type VideosMethod struct {
	client *Client
}

// ID returns information about the specified video
// https://dev.twitch.tv/docs/v5/reference/videos/#get-video
func (v *VideosMethod) ID(id uint) (*twitch.Video, error) {
	rel := fmt.Sprintf("videos/%d", id)

	video := new(twitch.Video)
	_, err := v.client.Get(rel, video)
	return video, err
}

// Top returns a list of the top viewed videos
// https://dev.twitch.tv/docs/v5/reference/videos/#get-top-videos
func (v *VideosMethod) Top(opt *twitch.ListOptions) (*twitch.TopVideosList, error) {
	rel := "videos/top"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	videos := new(twitch.TopVideosList)
	_, err := v.client.Get(rel, videos)
	return videos, err
}

// Followed returns a list of videos from channels the authenticated user is following
// https://dev.twitch.tv/docs/v5/reference/videos/#get-followed-videos
func (v *VideosMethod) Followed(opt *twitch.ListOptions) (*twitch.VideosList, error) {
	rel := "videos/followed"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	videos := new(twitch.VideosList)
	_, err := v.client.Get(rel, videos)
	return videos, err
}
