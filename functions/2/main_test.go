package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromedio(t *testing.T) {
	//Arrange
	notas := []float64{10, 8, 3}
	expectedResult := 7.0

	//Act
	resultado, err := promedio(notas...)
	if err != nil {
		t.Fatalf("error")
	}

	//Assert
	assert.Equal(t, expectedResult, resultado)
}

func TestPromedio_negativeNumber(t *testing.T) {
	//Arrange
	notas := []float64{-3, 2, 5}
	expectedResult := 0.0

	//Act
	resultado, err := promedio(notas...)
	if err == nil {
		t.Fatalf("error")
	}

	//Assert
	assert.Equal(t, expectedResult, resultado)

}

func TestPromedio_noInput(t *testing.T) {
	//Arrange
	notas := []float64{0.0}
	expectedResult := 0.0

	//Act
	resultado, err := promedio(notas...)
	if err != nil {
		t.Fatalf("error")
	}
	//Assert
	assert.Equal(t, expectedResult, resultado)
}
