package main

import (
	"fmt"
	"strings"
)

var coins = []string{}

func addCoin() {
	var input string

	fmt.Print("Enter a pairing (no spaces or dashes): ")
	fmt.Scanln(&input)
	pairing := strings.TrimSpace(input)
	coin := strings.ToUpper(pairing)

	test := getPrice(coin)
	if test == 0 {
		fmt.Println("Wrong input, try again.")
		addCoin()
	} else {
		coins = append(coins, coin)
		prompt()
	}
}

func indexOf(shitCoin string, coins []string) int {
	for index, coin := range coins {
		if shitCoin == coin {
			return index
		}
	}
	return -1
}

func deleteCoin() {
	var input string

	fmt.Print("Enter pairing to delete (no spaces or dashes): ")
	fmt.Scanln(&input)
	pairing := strings.TrimSpace(input)
	coin := strings.ToUpper(pairing)
	index := indexOf(coin, coins)
	if index != -1 {
		coins = append(coins[:index], coins[index+1:]...)
		fmt.Println("Coin deleted.")
		prompt()
	} else {
		fmt.Println("Input does not match any coins, try again.")
		deleteCoin()
	}
}
