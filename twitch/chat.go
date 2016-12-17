package twitch

// ChatBadges is a structure representing the chat badges that apply to a channel
// https://dev.twitch.tv/docs/v5/reference/chat/#get-chat-badges-by-channel
type ChatBadges struct {
	Admin       *ChatBadge `json:"admin,omitempty"`
	Broadcaster *ChatBadge `json:"broadcaster,omitempty"`
	GlobalMod   *ChatBadge `json:"global_mod,omitempty"`
	Mod         *ChatBadge `json:"mod,omitempty"`
	Staff       *ChatBadge `json:"staff,omitempty"`
	Subscriber  *ChatBadge `json:"subscriber,omitempty"`
	Turbo       *ChatBadge `json:"turbo,omitempty"`
}

// ChatBadge is an element of ChatBadges
// https://dev.twitch.tv/docs/v5/reference/chat/#get-chat-badges-by-channel
type ChatBadge struct {
	Alpha string `json:"alpha,omitempty"`
	Image string `json:"image,omitempty"`
	SVG   string `json:"svg,omitempty"`
}

// EmoticonSets is a structure representing a map of lists of EmoticonSet
// https://dev.twitch.tv/docs/v5/reference/chat/#get-chat-emoticons-by-set
type EmoticonSets struct {
	Sets map[string][]EmoticonSet `json:"emoticon_sets,omitempty"`
}

// EmoticonSet is an entry of EmoticonSets
// https://dev.twitch.tv/docs/v5/reference/chat/#get-chat-emoticons-by-set
type EmoticonSet struct {
	Code string `json:"code,omitempty"`
	ID   int    `json:"id,omitempty"`
}

// EmoticonsList is a structure representing a list of Emoticon
// https://dev.twitch.tv/docs/v5/reference/chat/#get-all-chat-emoticons
type EmoticonsList struct {
	Emoticons []Emoticon `json:"emoticons,omitempty"`
}

// Emoticon is an entry of EmoticonsList
// https://dev.twitch.tv/docs/v5/reference/chat/#get-all-chat-emoticons
type Emoticon struct {
	Regex  string          `json:"regex,omitempty"`
	Images []EmoticonImage `json:"images,omitempty"`
}

// EmoticonImage represents the image portion of Emoticon
// https://dev.twitch.tv/docs/v5/reference/chat/#get-all-chat-emoticons
type EmoticonImage struct {
	Width       int    `json:"width,omitempty"`
	Height      int    `json:"height,omitempty"`
	URL         string `json:"url,omitempty"`
	EmoticonSet int    `json:"emoticon_set,omitempty"`
}
