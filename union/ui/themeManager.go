package ui

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

//Move this section to themes manager
func applyStyle(dark bool) {
	css, err := gtk.CssProviderNew()
	if err != nil {
		log.Println("0 Failed to apply css style: ", err.Error())
		return
	}
	if dark {
		err = css.LoadFromPath("ui/dark.css")
		if err != nil {
			log.Println("1 Failed to apply css style: ", err.Error())
		}
	} else {
		err = css.LoadFromData("ui/light.css")
		if err != nil {
			log.Println("2 Failed to apply css style: ", err.Error())
		}
	}
}