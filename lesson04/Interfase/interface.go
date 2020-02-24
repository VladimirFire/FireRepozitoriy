package main

import (
	"fmt"
)

type man struct { // структура 1
	age float64
}

type woman struct { //структура 2
	age float64
}

func (s man) LenLife() float64 { // метод для структуры 1, считает недожитие мужчин до 100 лет
	return 100 - s.age
}
func (s woman) LenLife() float64 { // метод для структуры 2, считает недожитие женщин до 100 лет
	return 100 - s.age
}

type LenLifes interface { //интерфейс, объединяющий методы
	LenLife() float64
}

func SummALenLifes(pipls ...LenLifes) float64 { //функция, использующая интерфейс в роли входных данных
	res := 0.0
	for _, pipl := range pipls {
		res += pipl.LenLife()
	}
	return res
}

func main() {
	man1 := man{72} // задаем входные данные - продолжительность жизни людей
	man2 := man{82}
	man3 := woman{91}

	total := SummALenLifes(man1, man2, man3) / 3 //считаем среднюю нехватку лет для долгожительства
	fmt.Print("Люди, дожившие до 100 лет считаются долгожителями. Примерно столько лет не хватает людям моего окружения для дожития до долгожительства: ")
	fmt.Printf("%.2f", total)
}
