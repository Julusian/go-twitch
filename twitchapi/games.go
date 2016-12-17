package twitchapi

import (
	"github.com/julusian/go-twitch/twitch"
	"github.com/google/go-querystring/query"
)

// GamesMethod wraps up game api calls
// https://dev.twitch.tv/docs/v5/reference/games/
type GamesMethod struct {
	client *Client
}

// Top returns a list of games objects sorted by number of current viewers, most popular first
// https://dev.twitch.tv/docs/v5/reference/games/#get-top-games
func (g *GamesMethod) Top(opt *twitch.ListOptions) (*twitch.TopGamesList, error) {
	rel := "games/top"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	games := new(twitch.TopGamesList)
	_, err := g.client.Get(rel, games)
	return games, err
}
