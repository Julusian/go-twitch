package twitchapi

import (
	"fmt"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"
	"github.com/google/go-querystring/query"
)

type StreamsMethod struct {
	client *Client
}

// Returns a stream object if live.
func (s *StreamsMethod) Channel(id uint) (*twitch.SChannelS, error) {
	rel := fmt.Sprintf("streams/%d", id)

	stream := new(twitch.SChannelS)
	_, err := s.client.Get(rel, stream)
	return stream, err
}

// Returns a list of stream objects according to optional parameters.
func (s *StreamsMethod) List(opt *twitch.ListOptions) (*twitch.StreamsS, error) {
	rel := "streams"

	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	streams := new(twitch.StreamsS)
	_, err := s.client.Get(rel, streams)
	return streams, err
}

// Returns a list of featured (promoted) stream objects.
func (s *StreamsMethod) Featured(opt *twitch.ListOptions) (*twitch.FeaturedS, error) {
	rel := "streams/featured"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	featured := new(twitch.FeaturedS)
	_, err := s.client.Get(rel, featured)
	return featured, err
}

// Returns a summary of current streams.
func (s *StreamsMethod) Summary(opt *twitch.ListOptions) (*twitch.SummaryS, error) {
	rel := "streams/summary"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	summary := new(twitch.SummaryS)
	_, err := s.client.Get(rel, summary)
	return summary, err
}

// Returns a list of stream objects that the authenticated user is following.
func (s *StreamsMethod) followed(opt *twitch.ListOptions) (*twitch.FollowedS, error) {
	rel := "streams/followed"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	followed := new(twitch.FollowedS)
	_, err := s.client.Get(rel, followed)
	return followed, err
}
