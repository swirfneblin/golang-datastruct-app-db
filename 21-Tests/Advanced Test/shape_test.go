package shape

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulos", func(t *testing.T) {

		ret := Retangulo{10, 12}
		areaEsperada := 120.0
		areaRecebida := ret.Area()

		if areaRecebida != areaEsperada {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculos", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaRecebida != areaEsperada {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

}
