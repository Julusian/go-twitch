package twitchapi

import (
	"os"
	"testing"

	"github.com/julusian/go-twitch/twitch"
)

func TestUsersName(t *testing.T) {
	tc := newTestClient(t)

	_, err := tc.Users.Name(testUser)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestUsersUser(t *testing.T) {
	tc := newTestClient(t)

	_, err := tc.Users.User(testUserID)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestUsersSelf(t *testing.T) {
	tc := newTestClient(t)
	tc.OAuthToken = os.Getenv("ACCESSTOKEN")

	_, err := tc.Users.User(0)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestUsersFollows(t *testing.T) {
	tc := newTestClient(t)

	opt := &twitch.ListOptions{
		Limit:  1,
		Offset: 0,
	}

	_, err := tc.Users.Follows(testUserID, opt)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestUsersFollow(t *testing.T) {
	tc := newTestClient(t)

	_, err := tc.Users.Follow(testChannel, testUserID)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestUsersBlocks(t *testing.T) {
	tc := newTestClient(t)
	tc.OAuthToken = os.Getenv("ACCESSTOKEN")

	opt := &twitch.ListOptions{
		Limit:  1,
		Offset: 0,
	}

	_, err := tc.Users.Blocks(testChannel, opt)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}
