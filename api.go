package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getlastClose(pair string) float64 {
	url := "https://api.binance.com/api/v3/klines?symbol=" + pair + "&interval=1h&startTime=" + getLastTS() + "&endTime=" + getPrevTS()

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	resStr := string(responseData)
	conv1 := strings.ReplaceAll(resStr, "\"", "")
	conv2 := strings.Split(conv1, ",")
	close, _ := strconv.ParseFloat(conv2[4], 64)
	return close
}

func getWeekClose(pair string) float64 {
	url := "https://api.binance.com/api/v3/klines?symbol=" + pair + "&interval=1h&startTime=" + getWeekTS1() + "&endTime=" + getWeekTS2()

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	resStr := string(responseData)
	conv1 := strings.ReplaceAll(resStr, "\"", "")
	conv2 := strings.Split(conv1, ",")
	close, _ := strconv.ParseFloat(conv2[4], 64)
	return close
}

func getDif(cur float64, prev float64, dp int) string {
	white := "\033[0m"
	red := "\033[31m"
	green := "\033[32m"

	dif := (cur - prev) / prev

	if dp == 3 {
		if dif > 0 {
			p := dif * 100
			st := fmt.Sprintf("%v+%0.3f%%", green, p)
			return (st + white)
		} else {
			p := dif * 100
			st := fmt.Sprintf("%v%0.3f%%", red, p)
			return (st + white)
		}
	} else {
		if dif > 0 {
			p := dif * 100
			st := fmt.Sprintf("%v+%0.2f%%", green, p)
			return (st + white)
		} else {
			p := dif * 100
			st := fmt.Sprintf("%v%0.2f%%", red, p)
			return (st + white)
		}
	}
}

func get24hrPercent(pair string) string {
	url := "https://api.binance.com/api/v3/ticker/24hr?symbol=" + pair

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	white := "\033[0m"
	red := "\033[31m"
	green := "\033[32m"

	resStr := string(responseData)
	conv1 := strings.ReplaceAll(resStr, "\"", "")
	conv2 := strings.Split(conv1, ",")
	pcpStr := strings.TrimPrefix(conv2[2], "priceChangePercent:")
	pcp, _ := strconv.ParseFloat(pcpStr, 64)
	if pcp > 0 {
		cpp := fmt.Sprintf("%v+%0.2f%%", green, pcp)
		return (cpp + white)
	} else {
		cpp := fmt.Sprintf("%v%0.2f%%", red, pcp)
		return (cpp + white)
	}
}

func getPrice(pair string) float64 {
	url := "https://api.binance.com/api/v3/ticker/price?symbol=" + pair

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	resStr := string(responseData)
	conv1 := strings.Trim(resStr, "{\"symbol\":\""+pair+"\",\"price\":\"")
	conv2 := strings.Trim(conv1, "\"}")
	price, _ := strconv.ParseFloat(conv2, 64)
	return price
}
