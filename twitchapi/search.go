package twitchapi

import (
	"fmt"
	"net/url"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"

	"github.com/google/go-querystring/query"
)

type SearchMethod struct {
	client *Client
}

func (s *SearchMethod) Streams(q string, opt *twitch.ListOptions) (*twitch.StreamsS, error) {
	rel := "search/streams?query=" + url.QueryEscape(q)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "&" + v.Encode()
	}

	search := new(twitch.StreamsS)
	_, err := s.client.Get(rel, search)
	return search, err
}

func (s *SearchMethod) Games(q string, opt *twitch.ListOptions) (*twitch.SGamesS, error) {
	rel := fmt.Sprintf("search/games?query=%s", url.QueryEscape(q))
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "&" + v.Encode()
	}

	search := new(twitch.SGamesS)
	_, err := s.client.Get(rel, search)
	return search, err
}

func (s *SearchMethod) Channels(q string, opt *twitch.ListOptions) (*twitch.SChannelsS, error) {
	rel := fmt.Sprintf("search/channels?query=%s", url.QueryEscape(q))
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "&" + v.Encode()
	}

	search := new(twitch.SChannelsS)
	_, err := s.client.Get(rel, search)
	return search, err
}
