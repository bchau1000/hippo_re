package logging

import (
	"log"
	"os"
)

var Logfile *os.File = Init()
var Info = log.New(Logfile, "INFO ", log.LstdFlags|log.Lshortfile)
var Warning = log.New(Logfile, "WARNING ", log.LstdFlags|log.Lshortfile)
var Error = log.New(Logfile, "ERROR ", log.LstdFlags|log.Lshortfile)
var Fatal = log.New(Logfile, "FATAL ", log.LstdFlags|log.Lshortfile)

func Init() *os.File {
	logFile, err := os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Fatal error encountered opening log file: %v", err)
	}

	return logFile
}
