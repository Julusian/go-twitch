package twitchapi

import (
	"testing"

	"github.com/julusian/go-twitch/twitch"
)

func TestTeamsList(t *testing.T) {
	tc := newTestClient(t)

	opt := &twitch.ListOptions{
		Limit:  1,
		Offset: 0,
	}

	_, err := tc.Teams.List(opt)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestTeamsTeam(t *testing.T) {
	tc := newTestClient(t)

	_, err := tc.Teams.Team("staff")
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}
