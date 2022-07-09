package main

import (
	log "github.com/sirupsen/logrus"
)

func main(){
	log.WithFields(log.Fields{
		"user": "admin",
		"password": "1234",
		"key": "abcdefg",
	}).Info("Some interesting info")

	log.Warn("This is a warning")
	log.Error("This is an error")
	log.Fatal("This is a fatal error")
}