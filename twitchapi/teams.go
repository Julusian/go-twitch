package twitchapi

import (
	"github.com/julusian/go-twitch/twitch"
	"github.com/google/go-querystring/query"
)

// TeamsMethod wraps up team api calls
// https://dev.twitch.tv/docs/v5/reference/teams/
type TeamsMethod struct {
	client *Client
}

// List returns a list of all the teams
// https://dev.twitch.tv/docs/v5/reference/teams/#get-all-teams
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
	_, err := t.client.GetKraken(rel, teams)
	return teams, err
}

// Team returns the specified team
// https://dev.twitch.tv/docs/v5/reference/teams/#get-team
func (t *TeamsMethod) Team(name string) (*twitch.Team, error) {
	rel := "teams/" + name

	team := new(twitch.Team)
	_, err := t.client.GetKraken(rel, team)
	return team, err
}
