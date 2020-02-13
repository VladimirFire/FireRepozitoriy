package main

import (
	"fmt"
	"math"
	"strconv"
)

var (
	katet1, katet2, S, perimetr, gipotenuza float64
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
	fmt.Println("Введите длину катетов треугольника:")
	fmt.Scan(&katet1, &katet2)
	gipotenuza = math.Sqrt(math.Pow(katet1, 2) + math.Pow(katet2, 2))
	perimetr = gipotenuza + katet1 + katet2
	S = katet1 * katet2 / 2
	fmt.Println("Площадь =", Round(S, 2), "Периметр =", Round(perimetr, 2), "Гипотенуза =", Round(gipotenuza, 2))
}
