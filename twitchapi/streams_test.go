package twitchapi

import (
	"testing"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"
)

func TestStreamsChannel(t *testing.T) {
	tc := newTestClient(t)

	_, err := tc.Streams.Channel(testChannel)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestStreamsList(t *testing.T) {
	tc := newTestClient(t)

	opt := &twitch.ListOptions{
		Game:    "DayZ",
		Channel: "LIRIK",
		Limit:   1,
		Offset:  1,
		Hls:     false,
	}

	_, err := tc.Streams.List(opt)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestStreamsFeatured(t *testing.T) {
	tc := newTestClient(t)

	opt := &twitch.ListOptions{
		Limit:  1,
		Offset: 1,
	}

	_, err := tc.Streams.Featured(opt)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestStreamsSummary(t *testing.T) {
	tc := newTestClient(t)

	opt := &twitch.ListOptions{
		Game: "DayZ",
	}

	_, err := tc.Streams.Summary(opt)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}
