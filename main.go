// 		<> - Crypto price/volume momentum tracker
// 		Copyright (C) 2021  George Tevdo

// 		This program is free software: you can redistribute it and/or modify
// 		it under the terms of the GNU Affero General Public License as published
// 		by the Free Software Foundation, either version 3 of the License, or
// 		(at your option) any later version.

// 		This program is distributed in the hope that it will be useful,
// 		but WITHOUT ANY WARRANTY; without even the implied warranty of
// 		MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// 		GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"time"

	tm "github.com/buger/goterm"
)

func main() {
	tm.Clear()

	for {
		ui := tm.NewBox(50|tm.PCT, 18, 0)

		fmt.Fprintln(ui, " Pair       Price        1hr       24hr      7d     ")
		fmt.Fprintln(ui, " -------------------------------------------------- ")
		fmt.Fprintln(ui, "                                                    ")
		for _, coin := range coins {
			lc := getlastClose(coin)
			wc := getWeekClose(coin)
			price := getPrice(coin)
			dpp := get24hrPercent(coin)
			hpp := getDif(price, lc, 3)
			wpp := getDif(price, wc, 2)
			fmt.Fprintf(ui, " %v  %v    %v    %v    %v \n", coin, price, hpp, dpp, wpp)
			fmt.Fprintln(ui, "                                                    ")
		}

		tm.Print(tm.MoveTo(ui.String(), 1|tm.PCT, 5|tm.PCT))

		tm.Flush()

		time.Sleep(time.Second * 3)

		// addCoin()
		// deleteCoin()
	}
}
