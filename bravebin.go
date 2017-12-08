package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"host": "localhost",
	}).Info("One brave binary!")
}
