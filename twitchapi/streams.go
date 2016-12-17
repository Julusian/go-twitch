package twitchapi

import (
	"fmt"

	"github.com/julusian/go-twitch/twitch"
	"github.com/google/go-querystring/query"
)

// StreamsMethod wraps up stream api calls
// https://dev.twitch.tv/docs/v5/reference/streams/
type StreamsMethod struct {
	client *Client
}

// Channel returns a stream object of the specified channel if they are live
// https://dev.twitch.tv/docs/v5/reference/streams/#get-stream-by-channel
func (s *StreamsMethod) Channel(id uint) (*twitch.ChannelStream, error) {
	rel := fmt.Sprintf("streams/%d", id)

	stream := new(twitch.ChannelStream)
	_, err := s.client.Get(rel, stream)
	return stream, err
}

// List returns a list of stream objects according to optional parameters.
// https://dev.twitch.tv/docs/v5/reference/streams/#get-all-streams
func (s *StreamsMethod) List(opt *twitch.ListOptions) (*twitch.StreamsList, error) {
	rel := "streams"

	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	streams := new(twitch.StreamsList)
	_, err := s.client.Get(rel, streams)
	return streams, err
}

// Followed returns a list of stream objects that the authenticated user is following.
// https://dev.twitch.tv/docs/v5/reference/streams/#get-followed-streams
func (s *StreamsMethod) Followed(opt *twitch.ListOptions) (*twitch.StreamsList, error) {
	rel := "streams/followed"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	followed := new(twitch.StreamsList)
	_, err := s.client.Get(rel, followed)
	return followed, err
}

// Featured returns a list of featured stream objects
// https://dev.twitch.tv/docs/v5/reference/streams/#get-featured-streams
func (s *StreamsMethod) Featured(opt *twitch.ListOptions) (*twitch.FeaturedStreamsList, error) {
	rel := "streams/featured"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	featured := new(twitch.FeaturedStreamsList)
	_, err := s.client.Get(rel, featured)
	return featured, err
}

// Summary returns a summary of all current streams
// https://dev.twitch.tv/docs/v5/reference/streams/#get-streams-summary
func (s *StreamsMethod) Summary(opt *twitch.ListOptions) (*twitch.StreamSummary, error) {
	rel := "streams/summary"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	summary := new(twitch.StreamSummary)
	_, err := s.client.Get(rel, summary)
	return summary, err
}
