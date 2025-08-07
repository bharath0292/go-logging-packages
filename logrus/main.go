package main

import log "github.com/sirupsen/logrus"

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{
		"action": "payment",
		"amount": 400,
	}).Info("payment processed")
}
