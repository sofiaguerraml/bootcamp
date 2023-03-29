package ej1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {

	t.Run("Salario por debajo de $50000", func(t *testing.T) {
		//Arragne
		var salary float64 = 40000
		var expectedResult float64 = 0
		//Act
		tax := calculateTax(salary)
		//Assert
		assert.Equal(t, expectedResult, tax, "El impuesto no es el esperado")

	})
	t.Run("Salario entre $50000 y $150000", func(t *testing.T) {
		//Arragne
		var salary float64 = 60000
		var expectedResult float64 = 10200
		//Act
		tax := calculateTax(salary)
		//Assert
		assert.Equal(t, expectedResult, tax, "El impuesto no es el esperado")

	})
	t.Run("Salario por encima de 150000", func(t *testing.T) {
		//Arragne
		var salary float64 = 160000
		var expectedResult float64 = 43200
		//Act
		tax := calculateTax(salary)
		//Assert
		assert.Equal(t, expectedResult, tax, "El impuesto no es el esperado")

	})
}
