package twitchapi

import (
	"git.julusian.co.uk/botofdork/twitch-api/twitch"
	"github.com/google/go-querystring/query"
)

type GamesMethod struct {
	client *Client
}

// Returns a list of games objects sorted by number of current viewers, most
// popular first.
func (g *GamesMethod) Top(opt *twitch.ListOptions) (*twitch.TopsS, error) {
	rel := "games/top"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	games := new(twitch.TopsS)
	_, err := g.client.Get(rel, games)
	return games, err
}
