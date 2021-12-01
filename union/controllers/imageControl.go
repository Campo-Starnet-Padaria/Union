package controllers

import (
	"fmt"
	"log"

	"com.github/FelipeAlafy/union/manager"
	"com.github/FelipeAlafy/union/osmanager"
	"github.com/gotk3/gotk3/gdk"
)

var imageLimit int
var currentImage int = 0

func InitImageController() *gdk.Pixbuf {
	reload()
	imageLimit = len(instances[currentClient].Fotos)
	currentImage = 0
	
	return ActualImage() 
}

func crop(image *gdk.Pixbuf) *gdk.Pixbuf {
	nImage, err := image.ScaleSimple(380, 380, gdk.INTERP_BILINEAR)
	if err != nil {
		log.Println("Crop image error, because: ", err.Error())
	}
	return nImage  
}

func ActualImage() *gdk.Pixbuf {
	reload()
	imageLimit = len(instances[currentClient].Fotos)
	if imageLimit <= 0 {
		return nil
	}
	image, _ := gdk.PixbufNewFromFile(getPath(instances[currentClient].Fotos[currentImage]))
	return crop(image)
}

func NextImage() *gdk.Pixbuf {
	reload()
	imageLimit = len(instances[currentClient].Fotos)
	if imageLimit > 0 {
		currentImage++
		if currentImage == imageLimit {
			currentImage--
		}
		image, err := gdk.PixbufNewFromFile(getPath(instances[currentClient].Fotos[currentImage])) 
		if err != nil {
			log.Println("Next image error, because: ", err.Error())
		}
		return crop(image)
	}
	return nil
}

func PreviousImage() *gdk.Pixbuf {
	reload()
	imageLimit = len(instances[currentClient].Fotos)
	if imageLimit > 0 {
		currentImage--
		if currentImage == -1 {
			currentImage++
		}
		image, err := gdk.PixbufNewFromFile(getPath(instances[currentClient].Fotos[currentImage]))
		if err != nil {
			log.Println("Next image error, because: ", err.Error())
		}
		return crop(image)
	}
	return nil
}

func getPath(fileName string) string {
	return fmt.Sprint("union", so, "instances", so, instances[currentClient].Nome, so, fileName)
}

func InsertNewImage(uri string) {
	var images []string
	instances[currentClient].InsertAttachment(uri)
	images = append(images, manager.SubstringAfterLast(uri, so))

	instances[currentClient].Fotos = append(instances[currentClient].Fotos, images...)
	
	osmanager.ReWriteInstanceData(instances[currentClient])
	InitImageController()
}