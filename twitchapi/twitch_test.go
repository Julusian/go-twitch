package twitchapi

import (
	"net/http"
	"os"
	"testing"
)

const testChannel = 92016198 // Botofdork
const testUserID = 79110861  // Julusian
const testUser = "julusian"

func newTestClient(t *testing.T) *Client {
	client, err := NewClient(&http.Client{}, os.Getenv("CLIENTID"))
	if err != nil {
		t.Errorf("Failed to create client, %s", err)
	}

	return client
}
