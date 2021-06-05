package main

import "testing"

func TestSum(t *testing.T) {
	somatorio := soma(5, 5)

	if somatorio != 10 {
		t.Error("Soma dos números deverá retornar 10.")
	}

	if somatorio == 10 {
		t.Log("Teste passado com sucesso.")
	}
}
