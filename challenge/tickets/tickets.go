package tickets

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID      string
	Name    string
	Email   string
	Country string
	Hour    string
	Price   string
}

var tickets []Ticket

func ReadFile() (err error) {
	archivo, err := os.Open("./tickets.csv")
	if err != nil {
		return err
	}
	registros, err := csv.NewReader(archivo).ReadAll()
	if err != nil {
		return err
	}
	for _, reg := range registros {
		tickets = append(tickets, Ticket{
			ID:      reg[0],
			Name:    reg[1],
			Email:   reg[2],
			Country: reg[3],
			Hour:    reg[4],
			Price:   reg[5],
		})

	}
	return nil
}

func GetTotalTickets(destination string) (total int, err error) {
	gente := 0
	for _, v := range tickets {
		if v.Country == destination {
			gente += 1
		}
	}
	return gente, nil

}

func GetMornings(time string) (cant int, err error) {
	for _, v := range tickets {
		hora, err := strconv.Atoi(strings.Split(v.Hour, ":")[0])
		if err != nil {
			return 0, err
		}
		switch time {
		case "madrugada":
			if hora >= 0 && hora <= 6 {
				cant += 1
			}
		case "mañana":
			if hora >= 7 && hora <= 12 {
				cant += 1
			}
		case "tarde":
			if hora >= 13 && hora <= 19 {
				cant += 1
			}
		case "noche":
			if hora >= 20 && hora <= 23 {
				cant += 1
			}
		}

	}
	return cant, nil
}

func AverageDestination(destination string) (resultado float64, err error) {
	gente := 0
	for _, v := range tickets {
		if v.Country == destination {
			gente += 1
			resultado = (float64(gente) / float64(len(tickets))) * 100
		}

	}
	return resultado, nil
}
