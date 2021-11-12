package ui

import (
	"log"

	"com.github/FelipeAlafy/union/controllers"
	"github.com/gotk3/gotk3/gtk"
)

func photoView() *gtk.FileChooserButton {
	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	Err("photoView/WindowNew window", err)

	window.SetDefaultSize(400, 400)
	header, chooser, close := imagesHeaderbar() 
	window.SetTitlebar(header)
	close.Connect("clicked", func ()  {
		window.Hide()
	})

	hbox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	Err("photoView/BoxNew hbox", err)
	window.Add(hbox)

	//Buttons Forward And Backward
	forward, err := gtk.ButtonNewFromIconName("forward-symbolic", gtk.ICON_SIZE_BUTTON)
	Err("photoView/ButtonNew forward", err)

	backward, err := gtk.ButtonNewFromIconName("forward-rtl-symbolic", gtk.ICON_SIZE_BUTTON)
	Err("photoView/ButtonNew backward", err)

	forward.SetOpacity(0.10)
	backward.SetOpacity(0.10)
	
	image, _ := gtk.ImageNewFromPixbuf(controllers.InitImageController())

	window.Connect("activated", func ()  {
		image.SetFromPixbuf(controllers.InitImageController())
	})

	backward.Connect("clicked", func ()  {		 
		pixbuf := controllers.PreviousImage()
		image.SetFromPixbuf(pixbuf)
	})

	forward.Connect("clicked", func() { 	
		pixbuf := controllers.NextImage()
		image.SetFromPixbuf(pixbuf)
	})

	chooser.Connect("file-set", func ()  {
		uri := chooser.GetFilename()
		if err != nil {
			log.Fatal("FileChooser in photosUi, err: ", err.Error())
		}
		controllers.InsertNewImage(uri)
		chooser.SetFilename("")
		dialog := gtk.MessageDialogNew(window, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_CLOSE, "Imagem salva com êxito.")
		res := dialog.Run()
		if res == -7 {
			dialog.Close()
		} else {
			dialog.Close()
		}
	})

	hbox.PackStart(backward, false, true, 2)
	hbox.PackEnd(forward, false, true, 2)
	hbox.PackEnd(image, true, true, 2)
	window.ShowAll()
	return chooser
}