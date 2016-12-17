package twitch

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

type SGamesS struct {
	Games []GameS `json:"games,omitempty"`
}

type SChannelsS struct {
	Total    int        `json:"_total,omitempty"`
	Channels []ChannelS `json:"channels,omitempty"`
}

type SearchMethod struct {
	client *Client
}

func (s *SearchMethod) Streams(q string, opt *ListOptions) (*StreamsS, error) {
	rel := "search/streams?query=" + url.QueryEscape(q)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "&" + v.Encode()
	}

	search := new(StreamsS)
	_, err := s.client.Get(rel, search)
	return search, err
}

func (s *SearchMethod) Games(q string, opt *ListOptions) (*SGamesS, error) {
	rel := fmt.Sprintf("search/games?query=%s", url.QueryEscape(q))
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "&" + v.Encode()
	}

	search := new(SGamesS)
	_, err := s.client.Get(rel, search)
	return search, err
}

func (s *SearchMethod) Channels(q string, opt *ListOptions) (*SChannelsS, error) {
	rel := fmt.Sprintf("search/channels?query=%s", url.QueryEscape(q))
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "&" + v.Encode()
	}

	search := new(SChannelsS)
	_, err := s.client.Get(rel, search)
	return search, err
}
