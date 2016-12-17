package twitchapi

import "git.julusian.co.uk/botofdork/twitch-api/twitch"

type IngestsMethod struct {
	client *Client
}

func (i *IngestsMethod) List() (*twitch.IngestsList, error) {
	rel := "ingests"

	ingests := new(twitch.IngestsList)
	_, err := i.client.Get(rel, ingests)
	return ingests, err
}
