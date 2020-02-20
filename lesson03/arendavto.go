package main

import "fmt"

type param struct {
	model  string
	cvet   string
	god    int64
	probeg int64
	cena   float32
}

func main() {
	var s string
	car := make(map[string]param)
	car["Volvo"] = param{"xc90", "голубой", 2017, 62000, 1.34}
	car["Audi"] = param{"Q5", "Серый", 2018, 54000, 1.56}
	car["BMW"] = param{"X6", "черный", 2019, 34000, 3.12}
	fmt.Println("В наличии есть авто:")
	for name := range car {
		fmt.Println("-", name, car[name].model, "Цвет:", car[name].cvet, "Год выпуска:", car[name].god, "Пробег:", car[name].probeg, "Цена аренды (USD):", car[name].cena)
	}
	for {
		fmt.Println("Какую марку авто Вам выдать? (q - exit)")
		fmt.Scanln(&s)
		if s == "q" {
			break
		}
		delete(car, s)
		fmt.Println(s, "выдано!")
		fmt.Println("В наличии есть авто:")
		for name := range car {
			fmt.Println("-", name, car[name].model, "Цвет:", car[name].cvet, "Год выпуска:", car[name].god, "Пробег:", car[name].probeg, "Цена аренды (USD):", car[name].cena)
		}
	}
}
