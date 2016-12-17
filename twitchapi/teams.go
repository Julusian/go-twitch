package twitchapi

import (
	"git.julusian.co.uk/botofdork/twitch-api/twitch"
	"github.com/google/go-querystring/query"
)

type TeamsMethod struct {
	client *Client
}

func (t *TeamsMethod) List(opt *twitch.ListOptions) (*twitch.TeamsList, error) {
	rel := "teams"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	teams := new(twitch.TeamsList)
	_, err := t.client.Get(rel, teams)
	return teams, err
}

func (t *TeamsMethod) Team(name string) (*twitch.Team, error) {
	rel := "teams/" + name

	team := new(twitch.Team)
	_, err := t.client.Get(rel, team)
	return team, err
}
