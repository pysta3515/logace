package main

import (
	"github.com/sirupsen/logrus"
	log "github.com/pysta3515/logace"
)

func main() {
	// use a third log package
	log.SetLogEntity(logrus.New())

	log.Print("test Print:", 1, 2, 3)
	log.Println("test Println", 4, 5, 6)
	log.Printf("test Printf: %d, %d, %d", 7, 8, 9)

	log.Info("test Info:", 1, 2, 3)
	log.Infoln("test Infoln", 4, 5, 6)
	log.Infof("test Infof: %d, %d, %d", 7, 8, 9)

	log.Warn("test Warn:", 1, 2, 3)
	log.Warnln("test Warnln", 4, 5, 6)
	log.Warnf("test Warnf: %d, %d, %d", 7, 8, 9)
}
