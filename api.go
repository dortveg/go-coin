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

func getAveVol(pair string) float64 {
	url := "https://api.binance.com/api/v3/klines?symbol=" + pair + "&interval=1h&startTime=" + getOldTS() + "&endTime=" + getPrevTS()

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
	vol1, _ := strconv.ParseFloat(conv2[5], 64)
	vol2, _ := strconv.ParseFloat(conv2[17], 64)
	vol3, _ := strconv.ParseFloat(conv2[29], 64)
	aveVol := (vol1 + vol2 + vol3) / 3
	return aveVol
}

func getCurVol(pair string) float64 {
	url := "https://api.binance.com/api/v3/klines?symbol=" + pair + "&interval=1h&startTime=" + getPrevTS() + "&endTime=" + getCurTS()

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
	vol, _ := strconv.ParseFloat(conv2[5], 64)
	return vol
}

func get24hrPercent(pair string) float64 {
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

	resStr := string(responseData)
	conv1 := strings.ReplaceAll(resStr, "\"", "")
	conv2 := strings.Split(conv1, ",")
	pcpStr := strings.TrimPrefix(conv2[2], "priceChangePercent:")
	pcp, _ := strconv.ParseFloat(pcpStr, 64)
	return pcp
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
