package twitch

// IngestsList is a list of ingests available
// https://dev.twitch.tv/docs/v5/reference/ingests/#get-ingest-server-list
type IngestsList struct {
	Ingests []Ingest `json:"ingests,omitempty"`
}

// Ingest is an entry of IngestsList
// https://dev.twitch.tv/docs/v5/reference/ingests/#get-ingest-server-list
type Ingest struct {
	ID           int     `json:"_id,omitempty"`
	Availability float64 `json:"availability,omitempty"`
	Default      bool    `json:"default,omitempty"`
	Name         string  `json:"name,omitempty"`
	URLTemplate  string  `json:"url_template,omitempty"`
}
