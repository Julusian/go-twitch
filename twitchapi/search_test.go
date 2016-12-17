package twitchapi

import (
	"testing"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"
)

func TestSearchChannels(t *testing.T) {
	tc := newTestClient(t)

	opt := &twitch.ListOptions{
		Limit:  1,
		Offset: 0,
	}

	_, err := tc.Search.Channels("Star", opt)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestSearchStreams(t *testing.T) {
	tc := newTestClient(t)

	opt := &twitch.ListOptions{
		Limit:  1,
		Offset: 0,
	}

	_, err := tc.Search.Streams("Star", opt)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestSearchGames(t *testing.T) {
	tc := newTestClient(t)

	tru := true
	opt := &twitch.ListOptions{
		Live: &tru,
	}

	_, err := tc.Search.Games("Diablo", opt)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}
