// Streams methods of the twitch api.
// https://github.com/justintv/Twitch-API/blob/master/v3_resources/streams.md

package twitch

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

// used with GET /streams/:channel/
type SChannelS struct {
	Stream *StreamS `json:"stream,omitempty"`
}

// used with GET /streams
type StreamsS struct {
	Total   int       `json:"_total,omitempty"`
	Streams []StreamS `json:"streams,omitempty"`
}

// used with GET /streams/featured
type FeaturedS struct {
	Featured []FStreamS `json:"featured,omitempty"`
}

// used with GET /streams/summary
type SummaryS struct {
	Channels int `json:"channels,omitempty"`
	Viewers  int `json:"viewers,omitempty"`
}

// used with GET /streams/followed
type FollowedS struct {
	Total   int       `json:"_total,omitempty"`
	Streams []StreamS `json:"streams,omitempty"`
}

type StreamsMethod struct {
	client *Client
}

// Returns a stream object if live.
func (s *StreamsMethod) Channel(id uint) (*SChannelS, error) {
	rel := fmt.Sprintf("streams/%d", id)

	stream := new(SChannelS)
	_, err := s.client.Get(rel, stream)
	return stream, err
}

// Returns a list of stream objects according to optional parameters.
func (s *StreamsMethod) List(opt *ListOptions) (*StreamsS, error) {
	rel := "streams"

	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	streams := new(StreamsS)
	_, err := s.client.Get(rel, streams)
	return streams, err
}

// Returns a list of featured (promoted) stream objects.
func (s *StreamsMethod) Featured(opt *ListOptions) (*FeaturedS, error) {
	rel := "streams/featured"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	featured := new(FeaturedS)
	_, err := s.client.Get(rel, featured)
	return featured, err
}

// Returns a summary of current streams.
func (s *StreamsMethod) Summary(opt *ListOptions) (*SummaryS, error) {
	rel := "streams/summary"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	summary := new(SummaryS)
	_, err := s.client.Get(rel, summary)
	return summary, err
}

// Returns a list of stream objects that the authenticated user is following.
func (s *StreamsMethod) followed(opt *ListOptions) (*FollowedS, error) {
	rel := "streams/followed"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	followed := new(FollowedS)
	_, err := s.client.Get(rel, followed)
	return followed, err
}
