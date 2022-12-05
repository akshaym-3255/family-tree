package models

type Member struct {
	Name        string        `json:"name`
	Connections []Connections `json:connections`
}

type Connections struct {
	Relation string   `json:"relation`
	Ids      []string `json:"ids"`
}
