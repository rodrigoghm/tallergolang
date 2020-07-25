package util

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestPotencia(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	min := 2
	max := 8
	base := 0
	exp := 0
	for i := 1; i <= 1000; i++ {
		base = rand.Intn(max-min+1) + min
		exp = rand.Intn(max-min+1) + min
		fmt.Println("Prueba Nro ===> ", i)
		fmt.Println("=======================")
		fmt.Printf("base => [%d] || Exponente => [%d]\n\n", base, exp)
		if Potencia(uint64(base), uint64(exp)) != uint64(math.Pow(float64(base), float64(exp))) {
			t.Errorf("Error en prueba de testing - Valor esperado %v", math.Pow(float64(base), float64(exp)))
		}
	}
}
