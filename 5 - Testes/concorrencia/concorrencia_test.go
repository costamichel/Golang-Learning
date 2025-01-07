package main

import "testing"

func TestProcessarEmParalelo(t *testing.T) {
	dados := []int{0, 1, 2, 3, 4, 5}
	resultados := processarEmParalelo(dados)
	resultadosEsperados := []int{0, 2, 4, 6, 8, 10}

	if len(resultados) != len(resultadosEsperados) {
		t.Errorf("Eram esperados %d resultados mas foram obtidos %d resultados", len(resultadosEsperados), len(resultados))
	}

	for i, resultado := range resultados {
		if resultado != resultadosEsperados[i] {
			t.Errorf("Resultado incorreto para dado %d. Esperado %d e obtido %d", dados[i], resultadosEsperados[i], resultado)
		}
	}
}
