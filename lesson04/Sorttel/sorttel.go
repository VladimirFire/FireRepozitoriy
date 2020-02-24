package main

import (
	"fmt"
	"sort"
)

type tel struct { //тип, описывающий контакт в телефонной книге
	numb string
}

func main() {
	addressBook := make(map[string]tel) // создали массив контактов

	addressBook["Миша"] = tel{"78293467382"}
	addressBook["Никита"] = tel{"78293467382"}
	addressBook["Аня"] = tel{"78293467382"}
	addressBook["Катя"] = tel{"78293467382"}

	pipl := []string{}
	for name := range addressBook {
		pipl = append(pipl, name)
	}
	fmt.Println("Сортируем список людей и выводим их телефоны:")
	sort.Sort(sort.StringSlice(pipl)) //отсортировали список людей

	for i := range pipl {
		fmt.Println(pipl[i], addressBook[pipl[i]]) //выводим контакты по отсортированному списку
	}
}
