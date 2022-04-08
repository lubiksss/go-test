package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	//log.SetFormatter(&log.JSONFormatter{})
	//log.SetOutput(os.Stdout)
	//log.SetLevel(log.DebugLevel)
	//log.WithFields(log.Fields{
	//	"info": "Success",
	//}).Info("정상적으로 동작합니다.")

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.WithFields(log.Fields{
		"info": "Success",
	}).Info("정상적으로 동작합니다.")
	log.Info("test")
	log.Debug("test")
	log.Error("test")
	//log.Fatal("test")
}
