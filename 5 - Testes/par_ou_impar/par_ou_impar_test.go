package main

import (
	"fmt"
	"testing"
)

func TestParOuImpar(t *testing.T) {
	casos := []struct {
		entrada  int
		esperado string
	}{
		{entrada: 3, esperado: "ímpar"},
		{entrada: 0, esperado: "par"},
		{entrada: 102819287, esperado: "ímpar"},
		{entrada: -1, esperado: "ímpar"},
		{entrada: 3432742730, esperado: "par"},
	}

	for _, caso := range casos {
		t.Run(fmt.Sprintf("Teste %d", caso.entrada), func(t *testing.T) {
			resultado := ParOuImpar(caso.entrada)
			if resultado != caso.esperado {
				t.Errorf("Resultado obtido [%s] foi diferente do esperado [%s]", resultado, caso.esperado)
			}
		})
	}
}
