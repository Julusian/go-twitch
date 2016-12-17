package twitchapi

import "testing"

func TestIngestsList(t *testing.T) {
	tc := newTestClient(t)

	_, err := tc.Ingests.List()
	if err != nil {
		t.Errorf("error not nil: %s", err)
	}
}
