package main

import (
	"os"

	"github.com/apex/log"
	"github.com/sqrthree/debugfmt"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetHandler(debugfmt.New(os.Stdout))

	log.Debug("start")
	log.Info("start")
	log.Warn("start")
	log.Error("start")

	log.WithField("address", "http://localhost:3000").Debug("listening")
	log.WithField("address", "http://localhost:3000").Info("listening")
	log.WithField("address", "http://localhost:3000").Warn("listening")
	log.WithField("address", "http://localhost:3000").Error("listening")

	log.WithFields(log.Fields{
		"user": "foo",
		"file": "bar.png",
		"type": "image/png",
	}).Info("upload")

	log.WithFields(log.Fields{
		"user": "foo",
		"file": "bar.png",
		"type": "image/png",
	}).Warn("upload")

	log.WithFields(log.Fields{
		"user": "foo",
		"file": "bar.png",
		"type": "image/png",
	}).Error("upload")

	log.WithField("address", "http://localhost:3000").Fatal("listening")
}
