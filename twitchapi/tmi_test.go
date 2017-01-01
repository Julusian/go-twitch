package twitchapi

import "testing"

func TestHostsList(t *testing.T) {
	tc := newTestClient(t)

	_, err := tc.Tmi.Hosts(testChannel)
	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}
}
