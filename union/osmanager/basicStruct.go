package osmanager

import (
	"fmt"
	"log"
	"os"
	"runtime"
)
var so string

func BasicFolders() (error) {
	o := runtime.GOOS
    switch o {
    case "windows":
        fmt.Println("Windows")
		so = "\\"
    case "darwin":
        fmt.Println("MAC operating system")
		so = "/"
    case "linux":
        fmt.Println("Linux")
		so = "/"
    default:
        fmt.Printf("%s.\n", o)
		so = "/"
    }
	
	if _, err := os.Stat("union"); os.IsNotExist(err) {
		err := os.Mkdir("union", 0755)
		osPanic(err)
	} else {
		log.Println("union folder already exist skipping.")
	}
	if _, err := os.Stat("union/instances"); os.IsNotExist(err) {
		err := os.Mkdir(fmt.Sprint("union", so, "instances"), 0755)
		osPanic(err)
	} else {
		log.Println("union",so, "instances folder already exist skipping.")
	}
	return nil
}

func osPanic(e error) {
	if e != nil {
		panic(fmt.Sprint("union/osmanager/ - Error: ", e.Error()))
	}
}