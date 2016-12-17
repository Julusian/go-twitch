package twitch

import "github.com/google/go-querystring/query"

type TopsS struct {
	Total int    `json:"_total,omitempty"`
	Top   []TopS `json:"top,omitempty"`
}

type TopS struct {
	Channels int   `json:"channels,omitempty"`
	Viewers  int   `json:"viewers,omitempty"`
	Game     GameS `json:"game,omitempty"`
}

type GameS struct {
	ID          int      `json:"_id,omitempty"`
	Box         PreviewS `json:"box,omitempty"`
	GiantbombID int      `json:"giantbomb_id,omitempty"`
	Logo        PreviewS `json:"logo,omitempty"`
	Name        string   `json:"name,omitempty"`
	Popularity  int      `json:"popularity,omitempty"`
}

type GamesMethod struct {
	client *Client
}

// Returns a list of games objects sorted by number of current viewers, most
// popular first.
func (g *GamesMethod) Top(opt *ListOptions) (*TopsS, error) {
	rel := "games/top"
	if opt != nil {
		v, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		rel += "?" + v.Encode()
	}

	games := new(TopsS)
	_, err := g.client.Get(rel, games)
	return games, err
}
