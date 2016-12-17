package twitchapi

import (
	"os"
	"testing"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"
)

func TestVideosId(t *testing.T) {
	tc := newTestClient(t)

	_, err := tc.Videos.ID(44383045)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestVideosTop(t *testing.T) {
	tc := newTestClient(t)

	opt := &twitch.ListOptions{
		Limit:  1,
		Offset: 0,
		Game:   "Diablo",
		Period: "week",
	}

	_, err := tc.Videos.Top(opt)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestVideosFollowed(t *testing.T) {
	tc := newTestClient(t)
	tc.OAuthToken = os.Getenv("ACCESSTOKEN")

	opt := &twitch.ListOptions{
		Limit:  1,
		Offset: 0,
	}

	_, err := tc.Videos.Followed(opt)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}
