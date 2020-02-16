package main

import (
	"fmt"
	"math"
	"strconv"
)

const kurs = 63.15

var (
	rub float64
	usd float64
)

func Round(num float64, digits int) float64 {
	str := strconv.FormatFloat(float64(num), 'f', digits+1, 64)
	str, c := str[:len(str)-1], str[len(str)-1]
	f, _ := strconv.ParseFloat(str, 64)
	if c >= '5' && c <= '9' {
		f += 1 / math.Pow10(digits)
	}
	return f
}

func main() {
	fmt.Println("Сколько есть рублей?")
	fmt.Scan(&rub)
	usd = rub / kurs
	fmt.Println("Могу дать", Round(usd, 2), "USD")
}
