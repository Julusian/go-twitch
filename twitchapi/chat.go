package twitchapi

import (
	"fmt"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"
)

type ChatMethod struct {
	client *Client
}

func (c *ChatMethod) Badges(id uint) (*twitch.ChatBadgesS, error) {
	rel := fmt.Sprintf("chat/%d/badges", id)

	chatBadges := new(twitch.ChatBadgesS)
	_, err := c.client.Get(rel, chatBadges)
	return chatBadges, err
}

func (c *ChatMethod) Emoticons() (*twitch.EmoticonsS, error) {
	rel := "chat/emoticons"

	emoticons := new(twitch.EmoticonsS)
	_, err := c.client.Get(rel, emoticons)
	return emoticons, err
}

func (c *ChatMethod) EmoticonsBySet(set uint) (*twitch.EmoticonSetsS, error) {
	rel := "chat/emoticon_images"
	if set > 0 {
		rel += fmt.Sprintf("?emotesets=%d", set)
	}

	emoticons := new(twitch.EmoticonSetsS)
	_, err := c.client.Get(rel, emoticons)
	return emoticons, err
}
