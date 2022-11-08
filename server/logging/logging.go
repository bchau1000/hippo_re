package logging

import (
	"log"
	"os"
)

var Logfile, _ = os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
var Info = log.New(Logfile, "INFO ", log.LstdFlags|log.Lshortfile)
var Warning = log.New(Logfile, "WARNING ", log.LstdFlags|log.Lshortfile)
var Error = log.New(Logfile, "ERROR ", log.LstdFlags|log.Lshortfile)
var Debug = log.New(Logfile, "DEBUG ", log.LstdFlags|log.Lshortfile)
