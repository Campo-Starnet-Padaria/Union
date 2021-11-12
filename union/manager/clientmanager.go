package manager

import (
	"fmt"
	"io"
	"log"
	"os"
)

func GetClients(c []string) []Client {
	clients := make([]Client, len(c))
	for i, v := range c {
		json, err := JsonToClient([]byte(v))
		if err != nil {
			return nil
		}
		clients[i] = json	
	}
	return clients
}

func GetClient(c string) Client {
	json, err := JsonToClient([]byte(c))
	if err != nil {
		log.Fatal(err.Error())
	}
	return json
}

func (c Client) InsertAttachment(url string) error {
	fpath := SubstringAfterLast(url, "file://")
	file, _ := os.Open(fpath)
	path := fmt.Sprint("union/instances/", c.Nome, "/")
	filePath, err := makeFile(path, SubstringAfterLast(url, "/"))
	if err != nil {
		log.Println("I can't put ", SubstringAfterLast(url, "/"), " because ", err)
		return err
	}
	io.Copy(filePath, file)
	return nil
}

func makeFile(path, fileName string) (*os.File, error) {
	url := fmt.Sprint(path, fileName)
	file, err := os.Create(url)
	return file, err
}
