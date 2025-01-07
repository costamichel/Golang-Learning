package calculadora

import "testing"

func TestSoma(t *testing.T) {
	testes := []struct {
		nome     string
		a, b     int
		esperado int
	}{
		{"positivos", 4, 6, 10},
		{"negativos", -1, -3, -4},
		{"misto", -1, 3, 2},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := Soma(tt.a, tt.b)
			if resultado != tt.esperado {
				t.Errorf("Erro no teste %s: resultado %d é diferente do esperado %d", tt.nome, resultado, tt.esperado)
			}
		})
	}
}

func TestDivisao(t *testing.T) {
	testes := []struct {
		nome     string
		a, b     float64
		esperado float64
	}{
		{"divisaoPorZero", 5, 0, -1},
		{"positivos", 5, 2, 2.5},
		{"negativos", -3, -2, 1.5},
		{"misto", -3, 2, -1.5},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			resultado, err := Divisao(tt.a, tt.b)
			if tt.b == 0 {
				if err == nil {
					t.Errorf("Era esperado o erro de divisão por 0, mas não houve")
				}
			} else {
				if err != nil {
					t.Errorf("Houve um erro não esperado -> %s", err)
				} else {
					if resultado != tt.esperado {
						t.Errorf("Erro no teste '%s': resultado %f diferente do esperado %f", tt.nome, resultado, tt.esperado)
					}
				}
			}
		})
	}
}
