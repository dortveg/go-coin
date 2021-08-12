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
	//"strings"
)

func main() {

	// var btc coin = {
	// 	name: "BTCUSDT",
	// 	index: 0,
	// 	curPrice: 0,
	// 	prevPrice: 0,
	// 	curVol: 0,
	// 	prevVol: 0,
	// 	aveVolMin: 0,
	// 	oneMinTicks: []uint8{},
	// 	fiveMinTicks: []uint8{},
	// 	logs: []string{},
	// 	alertType: "none",
	// 	alertPrice: 0,
	// }

	nano := newCoin("NANOUSDT", 0)
	coins = append(coins, nano)

	fmt.Println(coins)
	// price := getPrice("ETHUSDT")
	// fmt.Println(price)
	// ave := getAveVol("ETHUSDT")
	// cur := getCurVol("ETHUSDT")
	// fmt.Printf("the ave is: %v, the cur is: %v", ave, cur)
	get24hrPercent("ETHUSDT")
	// fmt.Println("------------------")

	//addCoin()
}
