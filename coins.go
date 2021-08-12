package main

import (
	"fmt"
	"strings"
)

type coin struct {
	name         string
	index        int
	curPrice     float64
	prevPrice    float64
	curVol       float64
	prevVol      float64
	aveVolMin    float64
	oneMinTicks  []uint8
	fiveMinTicks []uint8
	logs         []string
	alertType    string
	alertPrice   float64
}

func newCoin(name string, index int) coin {
	pair := coin{
		name:         name,
		index:        index,
		curPrice:     0,
		prevPrice:    0,
		curVol:       0,
		prevVol:      0,
		aveVolMin:    0,
		oneMinTicks:  []uint8{},
		fiveMinTicks: []uint8{},
		logs:         []string{},
		alertType:    "none",
		alertPrice:   0,
	}
	return pair
}

var coins = []coin{}

func addCoin() {
	var index int
	var input string

	fmt.Print("Enter a pairing (no spaces or dashes): ")
	fmt.Scanln(&input)
	pairing := strings.TrimSpace(input)
	coin := strings.ToUpper(pairing)

	index = len(coins)

	newCoin := newCoin(coin, index)
	coins = append(coins, newCoin)

	fmt.Println(coins)
}
