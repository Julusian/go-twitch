package twitch

import "time"

type TeamsList struct {
	Teams []Team `json:"teams,omitempty"`
}

// Team object
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
