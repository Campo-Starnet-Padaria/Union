package manager

import (
	"fmt"
	"log"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func (c Client) GenProcuracaoCNPJ(city, representante, nacionalidade, estadoCivil, rgRepresentante, cpfRepresentante, numero, comp, rua, bairro, cidade, estado, cep string) {
	p := pdf.NewMaroto(consts.Portrait, consts.A4)
	p.SetPageMargins(20, 10, 20)

	buildHeader(p)
	c.bodyCnpj(p, representante, nacionalidade, estadoCivil, rgRepresentante, cpfRepresentante, numero, comp, rua, bairro, cidade, estado, cep)
	c.autoDate(p, city)

	err := p.OutputFileAndClose(fmt.Sprint("union/instances/", c.Nome, "/", c.Nome, " - Procuracao CNPJ.pdf"))
	if err != nil {
		log.Println(err.Error())
	}

}

func (c Client) bodyCnpj(m pdf.Maroto, representante, nacionalidade, estadoCivil, rgRepresentante, cpfRepresentante, numero, comp, rua, bairro, cidade, estado, cep string) {
	t1 := fmt.Sprint("Outorgante: ", c.Nome, " com sede à ", c.Rua, ", ", c.Numero, ", ", c.Bairro, ", ", c.Cep, ", ", c.Cidade, ", ")
	t2 := fmt.Sprint(c.Estado, ", ", "inscrita no CNPJ/MF nº", c.Cpf, " e inscrição estadual nº  ", c.Rg, ", ")
	t3 := fmt.Sprint("no estado de ", c.Estado,", representada por ", representante, ", ", nacionalidade, ", ", estadoCivil, ", ")  
	t4 := fmt.Sprint("portadora do R.G. Sob nº ", rgRepresentante, " expedido pela SSP-SE e com CPF nº ", cpfRepresentante) 
	t5 := fmt.Sprint("residente e domiciliado a ", rua, " nº ", numero, " ", comp, ", bairro ", bairro, ", ", cidade, ", ", estado, ", ")
	t6 := fmt.Sprint("CEP: ", cep, ", ", "que adquiriu um Sistema Fotovoltaico de, composto por módulos, inversor, instalação e homologação") 
	t7 :=  "na concessionária Cemig, pelo presente instrumento particular nomeia e constitui os engenheiros como seus"
	t8 := "Procuradores:"
	

	m.Row(5, func ()  {
		m.Col(100, func ()  {
			m.Text(t1, props.Text{
				Top: 1,
				Align:  consts.Left,
			})
		})
	})
	
	m.Row(5, func ()  {
		m.Col(100, func ()  {
			m.Text(t2, props.Text{
				Top: 1,
				Align:  consts.Left,
			})
		})
	})
	
	m.Row(5, func ()  {
		m.Col(100, func ()  {
			m.Text(t3, props.Text{
				Top: 1,
				Align:  consts.Left,
			})
		})
	})
	
	m.Row(5, func ()  {
		m.Col(100, func ()  {
			m.Text(t4, props.Text{
				Top: 1,
				Align:  consts.Left,
			})
		})
	})
	
	m.Row(5, func ()  {
		m.Col(100, func ()  {
			m.Text(t5, props.Text{
				Top: 1,
				Align:  consts.Left,
			})
		})
	})
	m.Row(5, func ()  {
		m.Col(100, func ()  {
			m.Text(t6, props.Text{
				Top: 1,
				Align:  consts.Left,
			})
		})
	})

	m.Row(5, func ()  {
		m.Col(100, func ()  {
			m.Text(t7, props.Text{
				Top: 1,
				Align:  consts.Left,
			})
		})
	})
	
	m.Row(5, func ()  {
		m.Col(100, func ()  {
			m.Text(t8, props.Text{
				Top: 1,
				Align:  consts.Left,
			})
		})
	})

	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text("Sr. Danilo Soares Costa, Solteiro, portador do CPF: 019.154.041-26, RG: 4971612,", props.Text{
				Top: 10,
			})
		})
	})

	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text("expedido por DGPC-GO, residente e domiciliado a Av Inglaterra, Quadra 117, Lote 1,", props.Text{
				Top: 3.5,
			})
		})
	})

	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text("Nº454, Setor Jardim Europa, Goiânia, Goiás, CEP: 74330-200, Telefone: (62) 99910-0030.", props.Text{
				Top: 3,
			})
		})
	})
	//Sergio dos Santos Marques
	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text("Sr. Sergio dos Santos Marques, Brasileiro, Casado, portador do CPF: 135.769.538-17, RG: 22.219.449-2,", props.Text{
				Top: 10,
			})
		})
	})

	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text("expedido por S-SP, residente e domiciliado a rua Doutor Almir Pinheiro Martins,", props.Text{
				Top: 3.5,
			})
		})
	})

	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text("Nº104, Campinas, Estado de São Paulo, CEP: 13624-060, Telefone: (11) 94034-0043", props.Text{
				Top: 3,
			})
		})
	})
	//Jobson Roberto Feitosa da Silva
	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text("Sr. Jobson Roberto Feitosa da Silva, Brasileiro, Solteiro, portador do CPF: 413.812.938-32,", props.Text{
				Top: 10,
			})
		})
	})

	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text("RG: 42.421.630-9, expedido por S-SP, residente e domiciliado a rua Doutor Almir Pinheiro Martins,", props.Text{
				Top: 3.5,
			})
		})
	})

	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text("Nº104, Campinas,  Estado de São Paulo, CEP: 13624-060, Telefone: (11) 94034-0043", props.Text{
				Top: 3,
			})
		})
	})
	//Rafael de Araujo Silva
	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text("Sr. Rafael de Araujo Silva, Brasileiro, Solteiro, portador do CPF: 356.670.848-80, RG: 44.887.188-9,", props.Text{
				Top: 10,
			})
		})
	})

	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text("expedido por S-SP, residente e domiciliado a rua Augusto Baer, Nº49, São Paulo,", props.Text{
				Top: 3.5,
			})
		})
	})

	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text("Estado de São Paulo, CEP: 02219-090, Telefone: (11) 94034-0043", props.Text{
				Top: 3,
			})
		})
	})

	//Default text
	m.Row(5, func ()  {
		m.Col(100, func ()  {
			m.Text("com poderes para atuar junto à Cemig, Concessionária de energia elétrica, relativo ao Processo", props.Text{
				Top: 20,
			})
		})
	})
	m.Row(5, func ()  {
		m.Col(100, func ()  {
			m.Text("de Homologação, Registro de acesso do Sistema fotovoltaico, podendo praticar qualquer ato, conferindo-lhe", props.Text{
				Top: 5,
			})
		})
	})
	m.Row(5, func ()  {
		m.Col(100, func ()  {
			m.Text("inclusive poderes para transigir, firmar acordo ou compromisso, formular e assinar TRT e ART e requerimentos,", props.Text{
				Top: 5,
			})
		})
	})
	m.Row(5, func ()  {
		m.Col(100, func ()  {
			m.Text("petições, ainda, substabelecer esta em outrem, com ou sem reservas de iguais poderes dando tudo por bom,", props.Text{
				Top: 5,
			})
		})
	})
	m.Row(5, func ()  {
		m.Col(100, func ()  {
			m.Text("firme e valioso, responsabilizando-se por todos os atos praticados no cumprimento deste instrumento, cessando os", props.Text{
				Top: 5,
			})
		})
	})
	m.Row(5, func ()  {
		m.Col(100, func ()  {
			m.Text("efeitos deste a partir da Solicitação de acesso da Cemig a ser emitido ou até ao prazo de 120 dias da emissão", props.Text{
				Top: 5,
			})
		})
	})
	m.Row(5, func ()  {
		m.Col(100, func ()  {
			m.Text("desta procuração.", props.Text{
				Top: 5,
			})
		})
	})
}