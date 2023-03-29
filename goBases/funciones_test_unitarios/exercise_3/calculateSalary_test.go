package ex3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSalary(t *testing.T) {
	t.Run("Categoria A", func(t *testing.T) {
		//Arragne
		minutes := 120.0
		category := "Categoria A"
		expectedSalary := 9000.0
		//Act
		salary := calculateSalary(minutes, category)
		//Assert
		assert.Equal(t, expectedSalary, salary, "El salario no es el esperado")

	})

	t.Run("Categoria B", func(t *testing.T) {
		//Arragne
		minutes := 120.0
		category := "Categoria B"
		expectedSalary := 3600.0
		//Act
		salary := calculateSalary(minutes, category)
		//Assert
		assert.Equal(t, expectedSalary, salary, "El salario no es el esperado")

	})

	t.Run("Categoria C", func(t *testing.T) {
		//Arragne
		minutes := 120.0
		category := "Categoria C"
		expectedSalary := 2000.0
		//Act
		salary := calculateSalary(minutes, category)
		//Assert
		assert.Equal(t, expectedSalary, salary, "El salario no es el esperado")

	})
}
