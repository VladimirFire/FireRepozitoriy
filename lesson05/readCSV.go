package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct {
	Firstname string   `json:"Имя"`
	Lastname  string   `json:"Фамилия"`
	Address   *Address `json:"Адрес"`
}

type Address struct {
	City  string `json:"Улица"`
	State string `json:"Дом"`
}

func main() {
	csvFile, _ := os.Open("file1.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var people []Person
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		people = append(people, Person{
			Firstname: line[0],
			Lastname:  line[1],
			Address: &Address{
				City:  line[2],
				State: line[3],
			},
		})
	}
	peopleJson, _ := json.Marshal(people)
	fmt.Println(string(peopleJson))
}
