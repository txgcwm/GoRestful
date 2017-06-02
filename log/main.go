package main

import (
	// "log"
	// "os"
	log "github.com/Sirupsen/logrus"
	"github.com/logmatic/logmatic-go"
)

// func log2file() {
// 	f,err :=os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
// 	if err !=nil{
// 		log.Fatal(err)
//   	}
 
//  	defer f.Close()
//  	log.SetOutput(f)
// 	log.Println("==========works==============")
// }

func main() {
	// use JSONFormatter
	log.SetFormatter(&logmatic.JSONFormatter{})

	contextLogger := log.WithFields(log.Fields{
		"common": "XXX common content XXX",
		"other": "YYY special context YYY",
	})

	contextLogger.Info("AAAAAAAAAAAA")
	contextLogger.Info("BBBBBBBBBBBB")

	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")

	// // log an event as usual with logrus
	// log.WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1 }).Info("My first ssl event from golang")
}