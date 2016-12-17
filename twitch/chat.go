// Streams methods of the twitch api.
// https://github.com/justintv/Twitch-API/blob/master/v3_resources/chat.md

package twitch

type ChatBadgesS struct {
	Admin       *ChatBadgeS `json:"admin,omitempty"`
	Broadcaster *ChatBadgeS `json:"broadcaster,omitempty"`
	GlobalMod   *ChatBadgeS `json:"global_mod,omitempty"`
	Mod         *ChatBadgeS `json:"mod,omitempty"`
	Staff       *ChatBadgeS `json:"staff,omitempty"`
	Subscriber  *ChatBadgeS `json:"subscriber,omitempty"`
	Turbo       *ChatBadgeS `json:"turbo,omitempty"`
}

type ChatBadgeS struct {
	Alpha string `json:"alpha,omitempty"`
	Image string `json:"image,omitempty"`
	SVG   string `json:"svg,omitempty"`
}

type EmoticonSetsS struct {
	Sets map[string]EmoticonSetS `json:"emoticon_sets,omitempty"`
}

type EmoticonSetS struct {
	Code string `json:"code,omitempty"`
	ID   int    `json:"id,omitempty"`
}

type EmoticonsS struct {
	Emoticons []EmoticonS `json:"emoticons,omitempty"`
}
type EmoticonS struct {
	Regex  string   `json:"regex,omitempty"`
	Images []ImageS `json:"images,omitempty"`
}

type ImageS struct {
	Width       int    `json:"width,omitempty"`
	Height      int    `json:"height,omitempty"`
	URL         string `json:"url,omitempty"`
	EmoticonSet int    `json:"emoticon_set,omitempty"`
}
