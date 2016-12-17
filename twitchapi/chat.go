package twitchapi

import (
	"fmt"

	"github.com/julusian/go-twitch/twitch"
)

// ChatMethod wraps up chat api calls
// https://dev.twitch.tv/docs/v5/reference/chat/
type ChatMethod struct {
	client *Client
}

// Badges returns a list of badges that apply to the specified channel
// https://dev.twitch.tv/docs/v5/reference/chat/#get-chat-badges-by-channel
func (c *ChatMethod) Badges(id uint) (*twitch.ChatBadges, error) {
	rel := fmt.Sprintf("chat/%d/badges", id)

	chatBadges := new(twitch.ChatBadges)
	_, err := c.client.Get(rel, chatBadges)
	return chatBadges, err
}

// EmoticonsBySet returns a list of emoticons available in the specified set
// https://dev.twitch.tv/docs/v5/reference/chat/#get-chat-emoticons-by-set
func (c *ChatMethod) EmoticonsBySet(set uint) (*twitch.EmoticonSets, error) {
	rel := "chat/emoticon_images"
	if set > 0 {
		rel += fmt.Sprintf("?emotesets=%d", set)
	}

	emoticons := new(twitch.EmoticonSets)
	_, err := c.client.Get(rel, emoticons)
	return emoticons, err
}

// Emoticons returns a list of emoticons available in chat
// https://dev.twitch.tv/docs/v5/reference/chat/#get-all-chat-emoticons
func (c *ChatMethod) Emoticons() (*twitch.EmoticonsList, error) {
	rel := "chat/emoticons"

	emoticons := new(twitch.EmoticonsList)
	_, err := c.client.Get(rel, emoticons)
	return emoticons, err
}
