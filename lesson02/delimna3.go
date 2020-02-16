package main

import (
	"fmt"
)

func main() {
    var N int
	fmt.Println("Введите число:")
	fmt.Scan(&N)
   
if N%3 == 0 {
        fmt.Println(N,"- делится на 3 БЕЗ остатка")
    } else {
        fmt.Println(N,"- делится на 3 с остатком")
    }
}
