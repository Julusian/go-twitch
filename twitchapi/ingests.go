package twitchapi

import "github.com/julusian/go-twitch/twitch"

// IngestsMethod wraps up ingest api calls
// https://dev.twitch.tv/docs/v5/reference/ingests/
type IngestsMethod struct {
	client *Client
}

// List returns a list of available ingest servers
// https://dev.twitch.tv/docs/v5/reference/ingests/#get-ingest-server-list
func (i *IngestsMethod) List() (*twitch.IngestsList, error) {
	rel := "ingests"

	ingests := new(twitch.IngestsList)
	_, err := i.client.Get(rel, ingests)
	return ingests, err
}
