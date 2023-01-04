package tickets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets(t *testing.T) {
	err := ReadFile("/Users/imihocevich/Desktop/actividad1/challenge/tickets.csv")
	if err != nil {
		t.Fatal(err)
	}
	//Arrange
	destination := "Brazil"
	expectedResult := 45
	//Act
	total, err := GetTotalTickets(destination)
	if err != nil {
		t.Fatalf("error")
	}
	//Assert
	assert.Equal(t, expectedResult, total)
}
