package twitch

type PreviewS struct {
	Large    string `json:"large,omitempty"`
	Medium   string `json:"medium,omitempty"`
	Small    string `json:"small,omitempty"`
	Template string `json:"template,omitempty"`
}

type ListOptions struct {
	Game       string `url:"game,omitempty"`
	Channel    string `url:"channel,omitempty"`
	Direction  string `url:"direction,omitempty"`
	Period     string `url:"period,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty"`
	Cursor     string `url:"cursor,omitempty"`
	Embeddable *bool  `url:"embeddable,omitempty"`
	Hls        bool   `url:"hls,omitempty"`
	Live       *bool  `url:"live,omitempty"`
	Cache      int64  `url:"_,omitempty"`
	Broadcasts bool   `url:"broadcasts,omitempty"`
}
