package manager

import (
	"encoding/json"
	"log"
	"os"
)

type Client struct {
	Nome string
	CPF string
	RG string
	Nascimento string
	Telefone string
	TelefoneFixo string
	ClienteCemig string
	InstalacaoCemig string
	EstadoCivil string
	//Address
	Numero int16
	Complemento int16
	Bairro string
	Cidade string
	Estado string //Default (MG)
	Cep string
	ZonaRural bool //Default false
	//attachments
	Fotos []os.File
	Observacao string
}

func (c Client) ClientToJson() ([]byte, error) {
	var jsonClient []byte
	log.Println("Converting ", c.Nome, " into Json")
	jsonClient, err := json.Marshal(c)
	log.Println("Conversion with error: ", err)
	return jsonClient, err
}

func JsonToClient(base []byte) (Client, error) {
	var client Client
	log.Println("Converting json into a Client")
	err := json.Unmarshal(base, &client)
	log.Println("Conversion with error: ", err)
	return client, err
}