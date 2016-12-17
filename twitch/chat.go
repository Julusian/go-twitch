// Streams methods of the twitch api.
// https://github.com/justintv/Twitch-API/blob/master/v3_resources/chat.md

package twitch

type ChatBadges struct {
	Admin       *ChatBadge `json:"admin,omitempty"`
	Broadcaster *ChatBadge `json:"broadcaster,omitempty"`
	GlobalMod   *ChatBadge `json:"global_mod,omitempty"`
	Mod         *ChatBadge `json:"mod,omitempty"`
	Staff       *ChatBadge `json:"staff,omitempty"`
	Subscriber  *ChatBadge `json:"subscriber,omitempty"`
	Turbo       *ChatBadge `json:"turbo,omitempty"`
}

type ChatBadge struct {
	Alpha string `json:"alpha,omitempty"`
	Image string `json:"image,omitempty"`
	SVG   string `json:"svg,omitempty"`
}

type EmoticonSets struct {
	Sets map[string]EmoticonSet `json:"emoticon_sets,omitempty"`
}

type EmoticonSet struct {
	Code string `json:"code,omitempty"`
	ID   int    `json:"id,omitempty"`
}

type EmoticonsList struct {
	Emoticons []Emoticon `json:"emoticons,omitempty"`
}
type Emoticon struct {
	Regex  string          `json:"regex,omitempty"`
	Images []EmoticonImage `json:"images,omitempty"`
}

type EmoticonImage struct {
	Width       int    `json:"width,omitempty"`
	Height      int    `json:"height,omitempty"`
	URL         string `json:"url,omitempty"`
	EmoticonSet int    `json:"emoticon_set,omitempty"`
}
