package main

import (
	"encoding/json"
	"io/ioutil"
	"k8s.io/klog/v2"
)

// this is implementation of option #1 of features to implement
// ideally, this would be a redis db which would keep data in memory
// and store data on disk
// this implementation is treated as interface - now it saves jsons to disk
// and it can be changed to work with redis/dynamodb/postgres/etc

const dbLocation = "localDb"

func SaveSnippet(snippet *Snippet) (err error) {
	file, err := json.MarshalIndent(*snippet, "", " ")
	if err != nil {
		klog.Error(err)
		return err
	}

	// snippet name is assumed to be unique
	//if err = ioutil.WriteFile(fmt.Sprintf("%s/%s.json", dbLocation, "test"), file, 0644); err != nil {
	klog.Info(*snippet)
	if err = ioutil.WriteFile(snippet.name, file, 0644); err != nil {
		klog.Error(err)
		return err
	}

	return nil
}

func ReadSnippet(snippetName string) (snippet *Snippet, err error) {

	//data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.json", dbLocation, snippetName))
	data, err := ioutil.ReadFile(snippetName)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	if err = json.Unmarshal(data, snippet); err != nil {
		klog.Error(err)
		return nil, err
	}

	return snippet, nil
}

// assume expiration time is extended by 30 seconds from existing expiry time

func UpdateSnippetExpiry(snippetName string) {

	snippet, err := ReadSnippet(snippetName)
	if err != nil {
		klog.Error(err)
		return
	}

	snippet.expires_at.Add(extensionExpiryDuration)

	if err := SaveSnippet(snippet); err != nil {
		klog.Error(err)
	}

}

func DeleteSnippet(snippetName string) {

}
