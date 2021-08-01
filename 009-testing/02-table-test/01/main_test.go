package main

import "testing"

func TestMiSuma(t *testing.T) {
	type prueba struct {
		datos     []int
		respuesta int
	}

	pruebas := []prueba{
		prueba{[]int{2, 4, 6}, 12},
		prueba{[]int{1, 5, 2}, 8},
		prueba{[]int{0, -1, 1}, 0},
		prueba{[]int{-10, 4, 20}, 14},
	}

	for _, x := range pruebas {
		v := miSuma(x.datos...)
		if v != x.respuesta {
			t.Error("Expected", x.respuesta, "Got", v)
		}
	}
}
