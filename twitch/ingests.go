package twitch

type IngestsList struct {
	Ingests []Ingest `json:"ingests,omitempty"`
}

type Ingest struct {
	ID           int     `json:"_id,omitempty"`
	Availability float64 `json:"availability,omitempty"`
	Default      bool    `json:"default,omitempty"`
	Name         string  `json:"name,omitempty"`
	URLTemplate  string  `json:"url_template,omitempty"`
}
