package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
	"net/http"
	"time"
)

const snippetsUrl = baseUrl + "/snippets"

// @ValidateInsertSnippetInput reads and validates input

func ValidateInsertSnippetInput(c *gin.Context, snippet *SnippetInput) (err error) {

	if err := c.Bind(snippet); err != nil {
		klog.Error(err)
		return err
	}

	// Additional checks go here

	return nil
}

func InsertSnippet(c *gin.Context) {

	snippetInput := new(SnippetInput)

	// Validate input
	if err := c.Bind(snippetInput); err != nil {
		klog.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Save full snippet url

	snippet := Snippet{
		Name:       snippetInput.Name,
		Expires_at: time.Now().Add(extensionExpiryDuration),
		Snippet:    snippetInput.Snippet,
		Url:        fmt.Sprintf("%s/%s.json", snippetsUrl, snippetInput.Name),
	}

	// Insert into db/save to disk

	if err := SaveSnippet(&snippet); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Return 201

	c.Status(http.StatusCreated)
}

func GetSnippet(c *gin.Context) {

	// Validate input

	snippetName := c.Param(snippet)

	// additional checks here
	if snippetName == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	snippet := new(Snippet)
	var err error

	if snippet, err = ReadSnippet(snippetName); err != nil {
		// If not exists or snippet expired, return 404
		if snippet != nil && snippet.Expires_at.Before(time.Now()) {
			DeleteSnippet(snippetName)
		}

		c.Status(http.StatusNotFound)
		return
	}

	if snippet != nil && snippet.Expires_at.Before(time.Now()) {
		DeleteSnippet(snippetName)
		c.Status(http.StatusNotFound)
		return
	}

	snippet.Expires_at = snippet.Expires_at.Add(extensionExpiryDuration)

	// add +30 seconds to expiration
	UpdateSnippetExpiry(snippet)

	// return 200 and snippet
	c.JSON(http.StatusOK, snippet)
}
