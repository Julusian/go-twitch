package twitch

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// PreviewImage holds a set of images for different sizes/resolutions
type PreviewImage struct {
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

type TwitchID struct {
	String  string
	Integer uint
}

func (this TwitchID) MarshalJSON() ([]byte, error) {
	str := this.String

	if this.Integer > 0 {
		str = fmt.Sprintf("%d", this.Integer)
	}

	if len(str) == 0 {
		return nil, fmt.Errorf("No id defined")
	}

	return []byte(str), nil
}

func (this *TwitchID) UnmarshalJSON(b []byte) error {
	var integer uint
	err := json.Unmarshal(b, &integer)
	if err == nil {
		// Treat it as an int!

		this.Integer = integer
		this.String = fmt.Sprintf("%d", integer)
		return nil
	}

	var str string
	err = json.Unmarshal(b, &str)
	if err == nil {
		if str == "" {
			return nil
		}

		integer, err := strconv.ParseUint(str, 10, 0)
		if err != nil {
			return err
		}

		this.Integer = uint(integer)
		this.String = str
		return nil
	}

	return err
}
