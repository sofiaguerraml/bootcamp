package ex2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	//Arrange
	var expectedResult float64 = 5
	//Act
	result := average(1, 2, 3, 4, 5, 6, 7, 8, 9)
	//Assert
	assert.Equal(t, expectedResult, result, "El promedio no es el esperado")

}
