package twitchapi

import (
	"fmt"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"
)

type ChatMethod struct {
	client *Client
}

func (c *ChatMethod) Badges(id uint) (*twitch.ChatBadges, error) {
	rel := fmt.Sprintf("chat/%d/badges", id)

	chatBadges := new(twitch.ChatBadges)
	_, err := c.client.Get(rel, chatBadges)
	return chatBadges, err
}

func (c *ChatMethod) Emoticons() (*twitch.EmoticonsList, error) {
	rel := "chat/emoticons"

	emoticons := new(twitch.EmoticonsList)
	_, err := c.client.Get(rel, emoticons)
	return emoticons, err
}

func (c *ChatMethod) EmoticonsBySet(set uint) (*twitch.EmoticonSets, error) {
	rel := "chat/emoticon_images"
	if set > 0 {
		rel += fmt.Sprintf("?emotesets=%d", set)
	}

	emoticons := new(twitch.EmoticonSets)
	_, err := c.client.Get(rel, emoticons)
	return emoticons, err
}
