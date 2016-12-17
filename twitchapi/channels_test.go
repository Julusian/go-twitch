package twitchapi

import (
	"os"
	"testing"

	"github.com/julusian/go-twitch/twitch"
)

func TestChannelsChannel(t *testing.T) {
	tc := newTestClient(t)

	_, err := tc.Channels.Channel(testChannel)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestChannelsSelf(t *testing.T) {
	tc := newTestClient(t)
	tc.OAuthToken = os.Getenv("ACCESSTOKEN")

	_, err := tc.Channels.Channel(0) // Logged in user
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestChannelsEditors(t *testing.T) {
	tc := newTestClient(t)
	tc.OAuthToken = os.Getenv("ACCESSTOKEN")

	_, err := tc.Channels.Editors(testChannel)
	if err != nil {
		t.Errorf("error not nil: %s", err)
	}
}

func TestChannelsFollows(t *testing.T) {
	tc := newTestClient(t)

	opt := &twitch.ListOptions{
		Limit:  1,
		Offset: 1,
	}

	_, err := tc.Channels.Follows(testChannel, opt)
	if err != nil {
		t.Errorf("error not nil: %s", err)
	}
}

// Note: cannot test due to test channel not having subscriptions enabled
// func TestChannelsSubscriptions(t *testing.T) {
// 	tc := newTestClient(t)
// 	tc.OAuthToken = os.Getenv("ACCESSTOKEN")

// 	opt := &twitch.ListOptions{
// 		Limit:  1,
// 		Offset: 1,
// 	}

// 	_, err := tc.Channels.Subscriptions(testChannel, opt)
// 	if err != nil {
// 		t.Errorf("error not nil: %+v", err)
// 	}
// }

func TestChannelsVideos(t *testing.T) {
	tc := newTestClient(t)

	opt := &twitch.ListOptions{
		Limit:  1,
		Offset: 1,
	}

	_, err := tc.Channels.Videos(testChannel, opt)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}
