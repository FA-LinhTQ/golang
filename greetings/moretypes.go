package main

import "fmt"

var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

type Animal struct {
	Name, Sciencename string
}

func (a Animal) getSciencename() string {
	return a.Sciencename
}

func getAnimalName(v Animal) string {
	return v.Name
}

func rangeFunc() {
	for _, v := range arr {
		fmt.Println(v)
	}
}

var animals map[string]Animal
var a = map[string]Animal{
	"Cat": {"Sphynx", "Sciencename"},
}

func main() {
	animals = make(map[string]Animal)
	animals["Cat"] = Animal{"Sphynx", "Sciencename"}
	// fmt.Println(animals["Cat"])
	delete(a, "Cat")
	a["Dog"] = Animal{"Doraemon", "Doremon"}
	v, ok := a["Cat"]
	fmt.Println("value: ", v, "Present? ", ok)
	fmt.Println(a)
	// rangeFunc()
}
