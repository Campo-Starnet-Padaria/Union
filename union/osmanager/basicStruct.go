package osmanager

import (
	"fmt"
	"os"
)

func BasicFolders() (error) {
	err := os.Mkdir("union", 0755)
	osPanic(err)
	err = os.Mkdir("union/instances", 0755)
	osPanic(err)
	return err
}

func osPanic(e error) {
	if e != nil {
		panic(fmt.Sprint("union/osmanager/basicStruct.go - Error: ", e))
	}
}