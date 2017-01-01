package twitchapi

import (
	"fmt"

	"github.com/julusian/go-twitch/twitch"
)

// TmiMethod wraps up tmi api calls
type TmiMethod struct {
	client *Client
}

// Hosts returns a list of the channels hosting the specified channel
func (t *TmiMethod) Hosts(channel uint64) (*twitch.TmiHostList, error) {
	rel := fmt.Sprintf("hosts?include_logins=1&target=%d", channel)

	hosts := new(twitch.TmiHostList)
	_, err := t.client.GetTMI(rel, hosts)
	return hosts, err
}
