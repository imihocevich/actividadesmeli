package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculadora(t *testing.T) {
	//Arrange
	mins := 0
	expectedResult := 0.00

	//Act
	salario := calculadora(mins, "A")

	//Assert
	assert.Equal(t, expectedResult, salario)
}
