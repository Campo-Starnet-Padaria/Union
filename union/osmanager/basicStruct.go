package osmanager

import (
	"fmt"
	"log"
	"os"
)

func BasicFolders() (error) {
	if _, err := os.Stat("union"); os.IsNotExist(err) {
		err := os.Mkdir("union", 0755)
		osPanic(err)
	} else {
		log.Println("union folder already exist skipping.")
	}
	if _, err := os.Stat("union/instances"); os.IsNotExist(err) {
		err := os.Mkdir("union/instances", 0755)
		osPanic(err)
	} else {
		log.Println("union/instances folder already exist skipping.")
	}
	return nil
}

func osPanic(e error) {
	if e != nil {
		panic(fmt.Sprint("union/osmanager/ - Error: ", e))
	}
}