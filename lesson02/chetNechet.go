package main

import (
	"fmt"
)

func main() {
    var N int
	fmt.Println("Введите число:")
	fmt.Scan(&N)
   
if N%2 == 0 {
        fmt.Println(N,"- четное число")
    } else {
        fmt.Println(N,"- нечетное число")
    }
}
