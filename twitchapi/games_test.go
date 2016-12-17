package twitchapi

import (
	"testing"

	"github.com/julusian/go-twitch/twitch"
)

func TestGamesTop(t *testing.T) {
	tc := newTestClient(t)

	opt := &twitch.ListOptions{
		Limit:  1,
		Offset: 0,
		Hls:    false,
	}

	_, err := tc.Games.Top(opt)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}
