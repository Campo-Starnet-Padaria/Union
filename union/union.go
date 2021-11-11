package main

import (
	"log"
	"os"

	"com.github/FelipeAlafy/union/osmanager"
	"com.github/FelipeAlafy/union/ui"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	osmanager.BasicFolders()
	app, err := gtk.ApplicationNew("com.github.FelipeAlafy.Union", glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal(err)
	}
	app.Connect("activate", func () {	ui.OnActivate(app)	})
	os.Exit(app.Run(os.Args))
}