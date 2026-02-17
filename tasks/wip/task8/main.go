// Собес: InfoWatch
// Задача:
// Что выведет функция

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type BidRequest struct {
	ID          int          `json:"id" bson:"id"`
	Impressions []Impression `json:"imp" bson:"imp"`
}

type Impression struct {
	BidFloor float64 `json:"bidfloor" bson:"bidfloor"`
}

func main() {
	req := BidRequest{
		ID: 1,
		Impressions: []Impression{
			{
				BidFloor: 0,
			},
		},
	}

	for range 3 {
		randomParameter := rand.Intn(10)
		go someAsyncFunc(req, randomParameter) // если рандом даст: 0, 2, 8
	}

	time.Sleep(time.Second)
}

func someAsyncFunc(req BidRequest, randomParameter int) {

	newImp := make([]Impression, len(req.Impressions))
	copy(newImp, req.Impressions)
	req.Impressions = newImp

	switch {
	case randomParameter == 0:
		req.Impressions[0].BidFloor = 0.1
	case randomParameter <= 5:
		req.Impressions[0].BidFloor = 8
	default:
		req.Impressions[0].BidFloor = 10
	}

	fmt.Println("Рандом -", randomParameter)
	sendNetworkRequest(req)
}

// что выведет принт на предложенный рандом
func sendNetworkRequest(req BidRequest) {
	fmt.Println(req.Impressions[0].BidFloor) // получим гонку данных, вывод будет рандомным 10 10 8 или 8 8 10 или 10 8 8 и т.д. В редок случае будет 0.1 8 10
}
