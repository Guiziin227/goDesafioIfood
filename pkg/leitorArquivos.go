package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LerArquivo(nomeArq, nomeBusca string) int {
	arquivo, err := os.Open(nomeArq)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo")
		os.Exit(1)
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)

	numeroLinha := 0
	ocorrencias := 0
	for scanner.Scan() {
		numeroLinha++
		linha := scanner.Text()

		if strings.Contains(strings.ToLower(linha), strings.ToLower(nomeBusca)) {
			ocorrencias++
			fmt.Printf("[%s] - linha %d: %s\n", nomeArq, numeroLinha, linha)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao scannear arquivo")
	}

	return ocorrencias
}
