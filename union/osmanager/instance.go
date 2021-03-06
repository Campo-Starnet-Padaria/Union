package osmanager

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"com.github/FelipeAlafy/union/manager"
)

//Attention to variable so, that defines which separator is used for correctly operating system. It was defined on basicStruct.

func CreateInstanceFolder(c manager.Client) error {
	path := "union" + so + "instances" + so + c.Nome
	if _, err := os.Stat(fmt.Sprint("union", so)); os.IsNotExist(err) {
		osPanic(err)
	} else if _, err := os.Stat(fmt.Sprint("union", so, "instances")); os.IsNotExist(err) {
		osPanic(err)
	} else if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(fmt.Sprint("union", so, "instances", so, c.Nome), 0755)
		osPanic(err)
	}
	return nil
}

func NewInstance(c manager.Client) *os.File {
	path := fmt.Sprint("union", so, "instances", so, c.Nome)
	file, err := os.Create(fmt.Sprint(path, so, c.Nome, ".json"))
	osPanic(err)
	return file
}

func InstanceData(bytes []byte, file *os.File) {
	_, err := file.WriteString(string(bytes))
	osPanic(err)
}

func ReWriteInstanceData(c manager.Client) {
	os.Remove(fmt.Sprint("union", so, "instances", so, c.Nome, so, c.Nome, ".json"))
	nFile := NewInstance(c)
	bytes, _ := c.ClientToJson()
	InstanceData(bytes, nFile)
}

//Get all instances
func GetInstances() []string {
	var jsons []string

	filepath.Walk(fmt.Sprint("union", so, "instances"), 
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				filename := strings.SplitAfterN(path, so, 4)
				if extension := strings.Split(filename[len(filename) - 1], "."); extension[1] == "json" {
					log.Println("Reading -> ", filename)
					if err != nil {
						log.Println(err)
					}
					dat, _ := os.ReadFile(path)
					jsons = append(jsons, string(dat))
				}
			}
			
			return nil
		},
	)
	return jsons
}

//Get instances by name
func GetSpecificInstances(specific string) string {
	var returnable string

	filepath.Walk(fmt.Sprint("union", so, "instances"), 
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				filename := strings.SplitAfterN(path, "/", 4)
				if extension := strings.Split(filename[len(filename) - 1], "."); extension[1] == "json" {
					log.Println("Reading -> ", filename)
					if  filename[len(filename) - 1] == fmt.Sprint(specific, ".json") {
						dat, _ := os.ReadFile(path)
						returnable = string(dat)
					}
					
				}
			}
			
			return nil
		},
	)
	return returnable
}