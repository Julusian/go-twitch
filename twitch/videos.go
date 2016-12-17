package twitch

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

type VVideosS struct {
	Total  int      `json:"_total,omitempty"`
	Videos []VideoS `json:"vods,omitempty"`
}

type VideosMethod struct {
	client *Client
}

func (v *VideosMethod) Id(id uint) (*VideoS, error) {
	rel := fmt.Sprintf("videos/%d", id)

	video := new(VideoS)
	_, err := v.client.Get(rel, video)
	return video, err
}

func (v *VideosMethod) Top(opt *ListOptions) (*VVideosS, error) {
	rel := "videos/top"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	videos := new(VVideosS)
	_, err := v.client.Get(rel, videos)
	return videos, err
}

func (v *VideosMethod) Followed(opt *ListOptions) (*VideosS, error) {
	rel := "videos/followed"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	videos := new(VideosS)
	_, err := v.client.Get(rel, videos)
	return videos, err
}
