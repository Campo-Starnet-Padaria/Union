package ui

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func headerbar () (*gtk.HeaderBar, []*gtk.Button) {
	hb, err := gtk.HeaderBarNew()
	if err != nil {
		log.Fatal("I cannot create headerbar because: ", err)
	}
	hb.SetShowCloseButton(true)
	hb.SetTitle("Union")
	new, edit, config, archive, folder, buttons := headerBarButtons()
	hb.PackStart(new)
	hb.PackStart(edit)
	hb.PackStart(folder)
	hb.PackEnd(config)
	hb.PackEnd(archive)
	return hb, buttons
}

func headerBarButtons() (*gtk.Button, *gtk.Button, *gtk.Button, *gtk.Button, *gtk.Button, []*gtk.Button) {
	//Add new, edit and config buttons
	//New
	newButton, err := gtk.ButtonNewWithLabel("Novo")
	if err != nil {
		log.Fatal("I cannot create button new on headerbar: ", err)
	}
	newButton.SetTooltipText("Cria um novo projeto")
	
	//Edit
	editButton, err := gtk.ButtonNewFromIconName("edit", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		log.Fatal("I cannot create button Edit on headerbar: ", err)
	}
	editButton.SetTooltipText("Permite editar o projeto que está na visualização.")

	//Config
	configButton, err := gtk.ButtonNewFromIconName("configuration", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		log.Fatal("I cannot create button Config on headerbar: ", err)
	}
	configButton.SetTooltipText("Ver as configurações rápidas.")

	//Archive project
	archiveButton, err := gtk.ButtonNewFromIconName("archive-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		log.Fatal("I cannot create button Config on headerbar: ", err)
	}
	archiveButton.SetTooltipText("Arquivar o projeto, quando feito não\npoderá mais ser visto.\nEntretanto, ainda está na pasta dele.")

	//Open Instance folder:
	instanceButton, err := gtk.ButtonNewFromIconName("folder", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		log.Fatal("I cannot create button Instance on headerbar: ", instanceButton)
	}
	instanceButton.SetTooltipText("Abre a pasta do projeto para você ver as fotos e documentos extras.")

	//array of buttons
	buttons := []*gtk.Button {newButton, editButton, configButton, archiveButton, instanceButton}

	return newButton, editButton, configButton, archiveButton, instanceButton, buttons
}