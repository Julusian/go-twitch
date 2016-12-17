package twitch

import "time"

// TeamsList is a list of Team
// https://dev.twitch.tv/docs/v5/reference/teams/#get-all-teams
type TeamsList struct {
	Teams []Team `json:"teams,omitempty"`
}

// Team is a structure representing a twitch team
// https://dev.twitch.tv/docs/v5/reference/teams/#get-all-teams
// https://dev.twitch.tv/docs/v5/reference/teams/#get-team
type Team struct {
	ID          int       `json:"_id,omitempty"`
	Background  string    `json:"background,omitempty"`
	Banner      string    `json:"banner,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	Info        string    `json:"info,omitempty"`
	Logo        string    `json:"logo,omitempty"`
	Name        string    `json:"name,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Users       []User    `json:"users,omitempty"`
}
