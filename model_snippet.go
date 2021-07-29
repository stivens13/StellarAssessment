package main

import "time"

const (
	snippet = "snippet"
	extensionExpiryDuration = 30 * time.Second
)

type SnippetInput struct {
	Name       string `json:"name,omitempty"`
	Expires_in int    `json:"expires_in,omitempty"`
	Snippet    string `json:"snippet,omitempty"`
}

type Snippet struct {
	Url        string    `json:"url,omitempty"`
	Name       string    `json:"name,omitempty"`
	Expires_at time.Time `json:"expires_at,omitempty"`
	Snippet    string    `json:"snippet,omitempty"`
}
