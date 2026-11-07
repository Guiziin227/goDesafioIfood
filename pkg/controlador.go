package pkg

import (
	"fmt"
	"sync"
	"time"
)

func Controle(arqs []string, nome string) {
	start := time.Now()

	var (
		wg sync.WaitGroup
	)

	wg.Go(func() {
		for i := range len(arqs) {

			ocorrencia := LerArquivo(arqs[i], nome)

			if ocorrencia != 0 {
				fmt.Println("RESUMO DA PESQUISA")
				fmt.Printf("[%s] - Ocorrências encontradas: %d\n", arqs[i], ocorrencia)
			}
		}
	})

	wg.Wait()
	fmt.Printf("Tempo total de busca foi de : %v\n", time.Since(start))
}
