package main

import "time"

const (
	snippet = "snippet"
	extensionExpiryDuration = 30 * time.Second
)

type SnippetInput struct {
	name string
	expires_in int
	snippet string
}

type Snippet struct {
	url string
	name string
	expires_at time.Time
	snippet string
}
