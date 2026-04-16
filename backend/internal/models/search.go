package models

type SearchResult struct {
	Source      string `json:"source"`
	SourceID    string `json:"sourceId"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ImageURL    string `json:"imageUrl"`
	ExternalURL string `json:"externalUrl"`
}
