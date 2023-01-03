package tickets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets(t *testing.T) {
	//Arrange
	destination := "Brazil"
	expectedResult := 45
	//Act
	total, err := GetTotalTickets(destination)
	if err != nil {
		t.Fatalf("error")
	}
	assert.Equal(t, expectedResult, total)
}
