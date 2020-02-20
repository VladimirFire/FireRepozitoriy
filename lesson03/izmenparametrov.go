package main

import (
	"fmt"
	"log"
	"strconv"
)

type param struct {
	model  string
	cvet   string
	god    int
	probeg int
	cena   float64
}

func main() {
	var s, s2, s3 string
	car := make(map[string]param)
	car["Volvo"] = param{"xc90", "голубой", 2017, 62000, 1.34}
	car["Audi"] = param{"Q5", "Серый", 2018, 54000, 1.56}
	car["BMW"] = param{"X6", "черный", 2019, 34000, 3.12}
	fmt.Println("В наличии есть авто:")
	for name := range car {
		fmt.Println("-", name, car[name].model, "Цвет:", car[name].cvet, "Год выпуска:", car[name].god, "Пробег:", car[name].probeg, "Цена аренды (USD):", car[name].cena)
	}
	for {
		fmt.Println("Какой марки авто хотите отредактировать? (q - exit)")
		fmt.Scanln(&s)
		if s == "q" {
			break
		}
		fmt.Println("Какой параметр хотите отредактировать? m - модель, c - цвет, y - год, p - пробег, a - цена (q - exit)")
		fmt.Scanln(&s2)
		if s2 == "q" {
			break
		}

		fmt.Println("Введите новое значение параметра (q - exit)")
		fmt.Scanln(&s3)
		if s3 == "q" {
			break
		}
		if s2 == "m" {
			fmt.Println(car[s].model)
		}

		switch s2 {
		case "m":
			car[s] = param{s3, car[s].cvet, car[s].god, car[s].probeg, car[s].cena}
		case "c":
			car[s] = param{car[s].model, s3, car[s].god, car[s].probeg, car[s].cena}
		case "y":
			{
				ints, err := strconv.Atoi(s3)
				if err != nil {
					log.Fatal(err)
				}
				car[s] = param{car[s].model, car[s].cvet, ints, car[s].probeg, car[s].cena}
			}
		case "p":
			{
				ints, err := strconv.Atoi(s3)
				if err != nil {
					log.Fatal(err)
				}
				car[s] = param{car[s].model, car[s].cvet, car[s].god, ints, car[s].cena}
			}
		case "a":
			{
				ints, err := strconv.ParseFloat(s3, 64)
				if err != nil {
					log.Fatal(err)
				}
				car[s] = param{car[s].model, car[s].cvet, car[s].god, car[s].probeg, ints}
			}
		default:
			fmt.Println("Неизвестный параметр")
		}
		fmt.Println("Отлично! Изменено!")
		fmt.Println("В наличии есть авто:")
		for name := range car {
			fmt.Println("-", name, car[name].model, "Цвет:", car[name].cvet, "Год выпуска:", car[name].god, "Пробег:", car[name].probeg, "Цена аренды (USD):", car[name].cena)
		}
	}
}
