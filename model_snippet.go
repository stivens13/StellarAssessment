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
	url string
	name string
	expires_at time.Time
	snippet string
}
