package manager

import (
	"encoding/json"
	"log"
)

type Client struct {
	Nome            string      `json:"Nome"`
	Cpf             string      `json:"CPF"`
	Rg              string      `json:"RG"`
	Nascimento      string      `json:"Nascimento"`
	Nacionalidade	string		`json:"Nacionalidade"`
	Telefone        string      `json:"Telefone"`
	TelefoneFixo    string      `json:"TelefoneFixo"`
	ClienteCemig    string      `json:"ClienteCemig"`
	InstalacaoCemig string      `json:"InstalacaoCemig"`
	EstadoCivil     string      `json:"EstadoCivil"`
	
	//Endereços
	Rua				string		`json:"Rua"`
	Numero          int         `json:"Numero"`
	Complemento     int         `json:"Complemento"`
	Bairro          string      `json:"Bairro"`
	Cidade          string      `json:"Cidade"`
	Estado          string      `json:"Estado"`
	Cep             string      `json:"Cep"`
	ZonaRural       bool        `json:"ZonaRural"`
	
	//Campo S&S utils
	ValorDoKit		float64		`json:"ValorDoKit"`
	ValorTotal		float64		`json:"ValorTotal"`
	Montagem		float64		`json:"Montagem"`
	Pago			bool		`json:"Pago"`
	Comissionario	string		`json:"Comissionario"`
	Comissao		float64		`json:"Comissao"`
	NotaKit			string		`json:"NotaKit"`
	NotaServico		string		`json:"NotaServico"`
	EntregaPlacas	string		`json:"EntregaPlacas"`
	Instalacao		string		`json:"Instalacao"`
	Finalizacao 	string		`json:"Finalizacao"`

	//CEMIG
	Rateio			[9]string	`json:"Rateio"`
	Geracao			string		`json:"Geracao"`
	MDGM			string		`json:"MDGM"` //Média de geração mês
	PKH				string		`json:"PKH"` //Preço Kw/hora

	//Anexos
	Fotos           []string	`json:"Fotos"`
	Observacao      string      `json:"Observacao"`

	//Unite utils
	Archived		bool 		`json:"Archived"`
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