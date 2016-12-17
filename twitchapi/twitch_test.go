package twitchapi

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"testing"

	"git.julusian.co.uk/botofdork/twitch-api/twitch"
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

type testStruct struct {
	ID twitch.TwitchID `json:"id"`
}

func testTwitchIdHelper(t *testing.T, src string, resInt uint, resStr string) {
	var res testStruct
	err := json.NewDecoder(strings.NewReader(src)).Decode(&res)
	if err != nil {
		t.Fatalf("Failed to parse: %s", err)
	}

	if res.ID.Integer != resInt {
		t.Errorf("Parsed wrong number! %d", res.ID.Integer)
	}
	if res.ID.String != resStr {
		t.Errorf("Parsed wrong string! %s", res.ID.String)
	}
}

func TestTwitchIDString(t *testing.T) {
	testTwitchIdHelper(t, `{"id":"123"}`, 123, "123")
}

func TestTwitchIDInt(t *testing.T) {
	testTwitchIdHelper(t, `{"id":123}`, 123, "123")
}

func TestTwitchIDEmptyString(t *testing.T) {
	testTwitchIdHelper(t, `{"id":""}`, 0, "")
}

func TestTwitchIDZeroInt(t *testing.T) {
	testTwitchIdHelper(t, `{"id":0}`, 0, "0")
}

func TestTwitchIDZeroStr(t *testing.T) {
	testTwitchIdHelper(t, `{"id":"0"}`, 0, "0")
}

func TestTwitchIDEmptyNegInt(t *testing.T) {
	var res testStruct
	err := json.NewDecoder(strings.NewReader(`{"id":-1}`)).Decode(&res)
	if err == nil {
		t.Fatalf("Expected unmarshal error")
	}
}

func TestTwitchIDEmptyNegStr(t *testing.T) {
	var res testStruct
	err := json.NewDecoder(strings.NewReader(`{"id":"-1"}`)).Decode(&res)
	if err == nil {
		t.Fatalf("Expected unmarshal error")
	}
}
