package main

import (
	"fmt"
	"./serverLogic"
	_ "./serverConnection"
)

func main() {
	//xml := "<resposta><retorno>0</retorno></resposta>"
	//xsd_path := "Arquivos/resposta.xsd"

	//serverLogic.ValidateXML(xml,xsd_path)
	//serverLogic.CheckXML(xml)

	resp := serverLogic.BuildXMLResponse("2")
	fmt.Println(resp)

}