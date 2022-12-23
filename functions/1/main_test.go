package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculadorasalario(t *testing.T) {
	//Arrange
	salario := 40000.0
	expectedResult := 0.0

	//Act
	resultado := calculadoraSalario(salario)

	//Assert
	assert.Equal(t, expectedResult, resultado)

}
