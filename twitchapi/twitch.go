package twitchapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// The root url of the twitch api
const rootURL = "https://api.twitch.tv/kraken/"

// Client is a struct containing methods to query the twitch api
type Client struct {
	client     *http.Client
	baseURL    *url.URL
	clientID   string
	OAuthToken string

	// Twitch api methods
	Channels *ChannelsMethod // https://dev.twitch.tv/docs/v5/reference/channels/
	Chat     *ChatMethod     // https://dev.twitch.tv/docs/v5/reference/chat/
	Games    *GamesMethod    // https://dev.twitch.tv/docs/v5/reference/games/
	Ingests  *IngestsMethod  // https://dev.twitch.tv/docs/v5/reference/ingests/
	Search   *SearchMethod   // https://dev.twitch.tv/docs/v5/reference/search/
	Streams  *StreamsMethod  // https://dev.twitch.tv/docs/v5/reference/streams/
	Teams    *TeamsMethod    // https://dev.twitch.tv/docs/v5/reference/teams/
	Users    *UsersMethod    // https://dev.twitch.tv/docs/v5/reference/users/
	Videos   *VideosMethod   // https://dev.twitch.tv/docs/v5/reference/videos/
}

// NewClient returns a new twitch client used to communicate with the API
// clientID is required to be able to query the api
func NewClient(httpClient *http.Client, clientID string) (*Client, error) {
	if len(clientID) == 0 {
		return nil, fmt.Errorf("ClientID is required to use the TwitchAPI")
	}

	baseURL, _ := url.Parse(rootURL)

	c := &Client{client: httpClient, baseURL: baseURL, clientID: clientID}
	c.Channels = &ChannelsMethod{client: c}
	c.Chat = &ChatMethod{client: c}
	c.Games = &GamesMethod{client: c}
	c.Ingests = &IngestsMethod{client: c}
	c.Search = &SearchMethod{client: c}
	c.Streams = &StreamsMethod{client: c}
	c.Teams = &TeamsMethod{client: c}
	c.Users = &UsersMethod{client: c}
	c.Videos = &VideosMethod{client: c}

	return c, nil
}

// Get issues an API get request and returns the API response. The response body is
// decoded and stored in the value pointed by r.
func (c *Client) Get(path string, r interface{}) (*http.Response, error) {
	// Try and parse url
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)

	// Create request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	req.Header.Add("Client-ID", c.clientID)

	// Add oauth token if supplied
	if len(c.OAuthToken) != 0 {
		req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", c.OAuthToken))
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotModified {
		return nil, fmt.Errorf("api error, response code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// try and parse response
	if r != nil {
		err = json.NewDecoder(resp.Body).Decode(r)
	}

	return resp, err
}
