package main

import (
	"desafioifood/pkg"
	"fmt"
)

func main() {

	arquivos := []string{"signatarios-out-2014.txt", "signatarios-fev-2002.txt", "signatarios-fev-2016.txt"}

	fmt.Println(" SISTEMA DE BUSCA DE NOMES ")

	fmt.Println("Digite o nome a ser buscado: ")

	var nome string
	fmt.Scanln(&nome)

	fmt.Printf("Buscando por: %s\n", nome)

	pkg.Controle(arquivos, nome)

}
