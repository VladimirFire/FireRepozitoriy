package main

import (
	"fmt"
)

var (
	vklad, procent, total float64
)



func main() {
	fmt.Println("Введите сумму вклада:")
	fmt.Scan(&vklad)
    fmt.Println("Введите ежегодный процент:")
	fmt.Scan(&procent)
    total=vklad+(vklad/100*procent)
     total+=(total/100*procent)
     total+=(total/100*procent)
     total+=(total/100*procent)
     total+=(total/100*procent)
    fmt.Println("Через 5 лет вы получите:",total)
}
