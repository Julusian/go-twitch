package twitchapi

import "testing"

func TestChatBadges(t *testing.T) {
	tc := newTestClient(t)

	_, err := tc.Chat.Badges(testChannel)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestChatEmoticonsBySet(t *testing.T) {
	tc := newTestClient(t)

	_, err := tc.Chat.EmoticonsBySet(19151)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}

func TestChatEmoticons(t *testing.T) {
	tc := newTestClient(t)

	_, err := tc.Chat.Emoticons()
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}
