package twitchapi

import (
	"fmt"
	"net/url"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"

	"github.com/google/go-querystring/query"
)

// SearchMethod wraps up search api calls
// https://dev.twitch.tv/docs/v5/reference/search/
type SearchMethod struct {
	client *Client
}

// Channels returns a list of channels that match the search query
// https://dev.twitch.tv/docs/v5/reference/search/#search-channels
func (s *SearchMethod) Channels(q string, opt *twitch.ListOptions) (*twitch.ChannelsSearch, error) {
	rel := fmt.Sprintf("search/channels?query=%s", url.QueryEscape(q))
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "&" + v.Encode()
	}

	search := new(twitch.ChannelsSearch)
	_, err := s.client.Get(rel, search)
	return search, err
}

// Streams returns a list of streams that match the search query
// https://dev.twitch.tv/docs/v5/reference/search/#search-streams
func (s *SearchMethod) Streams(q string, opt *twitch.ListOptions) (*twitch.StreamsList, error) {
	rel := "search/streams?query=" + url.QueryEscape(q)
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "&" + v.Encode()
	}

	search := new(twitch.StreamsList)
	_, err := s.client.Get(rel, search)
	return search, err
}

// Games returns a list of games that match the search query
// https://dev.twitch.tv/docs/v5/reference/search/#search-games
func (s *SearchMethod) Games(q string, opt *twitch.ListOptions) (*twitch.GamesSearch, error) {
	rel := fmt.Sprintf("search/games?query=%s", url.QueryEscape(q))
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "&" + v.Encode()
	}

	search := new(twitch.GamesSearch)
	_, err := s.client.Get(rel, search)
	return search, err
}
