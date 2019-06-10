package core

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	file, err := os.OpenFile("sparta.log", os.O_APPEND|os.O_CREATE, 666)
	if err != nil {
		log.Fatalln("fail to create sparta.log file!")
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Panic("")
		}
	}()
	Logger = log.New(file, "", log.LstdFlags|log.Lshortfile)

	Logger.SetFlags(log.LstdFlags | log.Lshortfile)
}