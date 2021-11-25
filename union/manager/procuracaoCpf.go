package manager

import (
	"fmt"
	"log"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func (c Client) GenProcuracaoCPF(city string) {

	p := pdf.NewMaroto(consts.Portrait, consts.A4)
	p.SetPageMargins(20, 10, 20)

	buildHeader(p)
	c.body(p)
	c.autoDate(p, city)
	

	err := p.OutputFileAndClose(fmt.Sprint("union/instances/", c.Nome, "/", c.Nome, " - Procuracao CPF.pdf"))
	if err != nil {
		log.Println(err.Error())
	}
}

func buildHeader(m pdf.Maroto) {
	m.Row(20, func ()  {
		m.Col(12, func ()  {
			m.Text("PROCURAÇÃO INSTRUMENTO PARTICULAR ", props.Text{
				Top: 3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: color.Color {
					Red: 88,
					Green: 80,
					Blue: 90,
				},
			})
		})
	})
}

func (c Client) body(m pdf.Maroto) {
	//Client start
	t1 := fmt.Sprint("Outorgante: ", c.Nome, ", ", c.Nacionalidade, ", ",  c.EstadoCivil, ", portadora do R.G. Sob nº ", c.Rg)
	t2 := fmt.Sprint("expedido pela SESP-PR e com CPF nº ", c.Cpf, " residente e domiciliado na rua ", c.Rua, ",")
	t3 := fmt.Sprint("nº ", c.Numero, " complemento: ", c.Complemento, " no bairro: ", c.Bairro," na cidade: ", c.Cidade)
	t4 := fmt.Sprint("do estado de Minas Gerais cep: ", c.Cep, " que adquiriu um Sistema Fotovoltaico e deseja fazer a homologação")
	t5 := 			 "na concessionaria Cemig, pelo presente instrumento particular nomeia e constitui os"
	t6 := 			 "engenheiros como seus Procuradores:" 

	m.Row(5, func() {
		m.Col(100, func() {
			m.Text(t1, props.Text{
				Top: 1,
				Align: consts.Left,
			})
		})
	})

	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text(t2, props.Text{
				Top: 1,
				Align: consts.Left,
			})
		})
	})
	
	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text(t3, props.Text{
				Top: 1,
				Align: consts.Left,
			})
		})
	})
	
	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text(t4, props.Text{
				Top: 1,
				Align: consts.Left,
			})
		})
	})
	
	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text(t5, props.Text{
				Top: 1,
				Align: consts.Left,
			})
		})
	})
	
	m.Row(5, func() {
		m.Col(100, func ()  {
			m.Text(t6, props.Text{
				Top: 1,
				Align: consts.Left,
			})
		})
	})
	//Client end - Engineers start
	//Danilo Soares Costa
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

func (c Client) autoDate(m pdf.Maroto, city string) {
	m.Row(10, func() {
		m.Col(12, func ()  {
			m.Text(fmt.Sprint(city, " - ", c.Estado, " ", dateFormat()), props.Text{
				Top: 10,
				Align: consts.Center,
			})
		})
	})

	m.Row(5, func ()  {
		m.Col(12, func ()  {
			m.Text("___________________________________________________________________", props.Text{
				Top: 100,
				Align: consts.Center,
			})
		})
	})

	m.Row(5, func ()  {
		m.Col(12, func ()  {
			m.Text("_________________________", props.Text{
				Top: 7,
				Align: consts.Center,
			})
		})
	})
}

func dateFormat() string {
	date := time.Now().AddDate(0, 0, 1)
	log.Println("Date: ", date)

	data := fmt.Sprint(date.Day(), " de ")
	
	if date.Month().String() ==  "January" {
		data = fmt.Sprint(data + "Janeiro de ")
	} else if date.Month().String() == "February" { 
		data = fmt.Sprint(data + "Fevereiro de ")
	} else if date.Month().String() == "March" {	
		data = fmt.Sprint(data + "Março de ")
	} else if  date.Month().String() == "April" {
		data = fmt.Sprint(data + "Abril de ")
	} else if date.Month().String() == "May" {
		data = fmt.Sprint(data + "Maio de ")
	} else if date.Month().String() =="June"{
		data = fmt.Sprint(data + "Junho de ")
	} else if date.Month().String() == "July" {
		data = fmt.Sprint(data + "Julho de ")
	} else if date.Month().String() == "August" {
		data = fmt.Sprint(data + "Agosto de ")
	} else if date.Month().String() == "September" {
		data = fmt.Sprint(data + "Setembro de ")
	} else if date.Month().String() =="October" {
		data = fmt.Sprint(data + "Outubro de ")
	} else if date.Month().String() =="November" {
		data = fmt.Sprint(data + "Novembro de ")
	} else if date.Month().String() =="December" {
		data = fmt.Sprint(data + "Dezembro de ")
	}

	data = fmt.Sprint(data, date.Year())
	return data
}