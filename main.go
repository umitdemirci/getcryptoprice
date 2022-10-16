package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const BINANCE_PRICE_URL = "https://api.binance.com/api/v3/ticker/price?symbol=%s%s"
const CURRENCY = "USDT"

type Crypto struct {
	Price float64 `json:"price,string"`
}

const USAGE = `
      getcryptoprice [parity]
	  Example usages:
		getcryptoprice BTC
		getcryptoprice ETH
		getcryptoprice CHZ
`

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Println(USAGE)
		os.Exit(1)
	}
	symbol := flag.Args()[0]
	symbol = strings.TrimSpace(symbol)
	symbol = strings.ToUpper(symbol)
	resp, err := http.Get(fmt.Sprintf(BINANCE_PRICE_URL, symbol, CURRENCY))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var crypto Crypto
	err = json.Unmarshal(body, &crypto)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s price is %.2f %s\n", symbol, crypto.Price, CURRENCY)
}
