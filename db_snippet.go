package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8s.io/klog/v2"
	"os"
)

// this is implementation of option #1 of features to implement
// ideally, this would be a redis db which would keep data in memory
// and store data on disk
// this implementation is treated as interface - now it saves jsons to disk
// and it can be changed to work with redis/dynamodb/postgres/etc

const dbLocation = "localDb"

func init() {
	if _, err := os.Stat(dbLocation); os.IsNotExist(err) {
		if err := os.MkdirAll(dbLocation, 0700); err != nil {
			klog.Error(err)
		}
	}
}

func SaveSnippet(snippet *Snippet) (err error) {
	file, err := json.Marshal(snippet)
	if err != nil {
		klog.Error(err)
		return err
	}

	// snippet name is assumed to be unique
	if err = ioutil.WriteFile(fmt.Sprintf("%s/%s.json", dbLocation, snippet.Name), file, 0644); err != nil {
		klog.Error(err)
		return err
	}

	return nil
}

func ReadSnippet(snippetName string) (*Snippet, error) {

	data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.json", dbLocation, snippetName))
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	snippet := &Snippet{}

	if err = json.Unmarshal(data, snippet); err != nil {
		klog.Error(err)
		return nil, err
	}

	return snippet, nil
}

// assume expiration time is extended by 30 seconds from existing expiry time

func UpdateSnippetExpiry(snippet *Snippet) {

	if err := SaveSnippet(snippet); err != nil {
		klog.Error(err)
	}

}

func DeleteSnippet(snippetName string) {
	if err := os.Remove(fmt.Sprintf("%s/%s.json", dbLocation, snippetName)); err != nil {
		klog.Info(err)
	}
}
