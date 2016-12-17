package twitch

type IngestsS struct {
	Ingests []IngestS `json:"ingests,omitempty"`
}

type IngestS struct {
	ID           int     `json:"_id,omitempty"`
	Availability float64 `json:"availability,omitempty"`
	Default      bool    `json:"default,omitempty"`
	Name         string  `json:"name,omitempty"`
	URLTemplate  string  `json:"url_template,omitempty"`
}

type IngestsMethod struct {
	client *Client
}

func (i *IngestsMethod) List() (*IngestsS, error) {
	rel := "ingests"

	ingests := new(IngestsS)
	_, err := i.client.Get(rel, ingests)
	return ingests, err
}
