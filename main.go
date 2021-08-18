package main

import (
	"fmt"
	"strings"
	"time"

	tm "github.com/buger/goterm"
)

func prompt() {
	var input string

	fmt.Print("Welcome. Type 'a' - add a coin, 's' - start, 'd' - delete a coin: ")
	fmt.Scanln(&input)
	trm := strings.TrimSpace(input)
	choice := strings.ToUpper(trm)

	switch choice {
	case "A":
		addCoin()
	case "S":
		startTracking()
	case "D":
		deleteCoin()
	default:
		fmt.Println("Invalid input.")
		prompt()
	}
}

func startTracking() {
	tm.Clear()

	for {
		ui := tm.NewBox(57, (len(coins)*2)+6, 0)

		fmt.Fprintln(ui, " PAIR       PRICE       1HR%       1D%       7D%     ")
		fmt.Fprintln(ui, " ___________________________________________________ ")
		fmt.Fprintln(ui, "                                                     ")
		for _, coin := range coins {
			lc := getlastClose(coin)
			wc := getWeekClose(coin)
			price := getPrice(coin)
			dpp := get24hrPercent(coin)
			hpp := getDif(price, lc, 3)
			wpp := getDif(price, wc, 2)
			fmt.Fprintf(ui, " %-8s   %-8s   %v    %v    %v \n", coin, fmt.Sprint(price), hpp, dpp, wpp)
			fmt.Fprintln(ui, "                                                    ")
		}
		fmt.Fprintf(ui, " Currently tracking %v coins. Press Ctrl+C to quit.", len(coins))

		tm.Print(tm.MoveTo(ui.String(), 1|tm.PCT, 5|tm.PCT))

		tm.Flush()

		time.Sleep(time.Second * 3)
	}
}

func main() {
	//test()
	prompt()
}
