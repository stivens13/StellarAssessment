package main

import (
	"k8s.io/klog/v2"
)

var (
	port = ":80"
)

func main() {

	klog.Info("Creating Assessment Server")

	router := CreateRouter()

	klog.Info("Starting Assessment Server")
	klog.Exit(router.Run(port))
}