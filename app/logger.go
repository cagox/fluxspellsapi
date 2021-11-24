package app

import (
	"log"
	"os"
)

//StartLogging() starts the logger.
func StartLogging() {
	//Set up logging
	var err error
	Config.LogFile, err = os.OpenFile(Config.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		panic("Couldn't Open Log File")
	}
	Config.Logger = log.New(Config.LogFile, "ocfsocial:", log.LstdFlags)
}

//StopLogging() stops the logger.
func StopLogging() {
	_ = Config.LogFile.Close()
	//TODO: Add error handling.
}
