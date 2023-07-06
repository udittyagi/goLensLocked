package main

import (
	"html/template"
	"os"
)

type Address struct {
	City string
	Pin int64
}
type Person struct {
	Name string
	Bio string
	Age int
	AddressMap map[string]Address
	Arr []int
	ArrAddress []Address
	IsLoading bool
	IsActive bool
}

func main() {
	t ,err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	p1 := Person {
		Name: "Udit", 
		Bio: `<script>alert("Hello There!!")</script>`,
		Age: 24,
		AddressMap: map[string]Address{
			"city1": {"Hapur", 245101},
			"city2": {"Meertt", 245304},
		},
		Arr: []int{1,2,3},
		ArrAddress: []Address{{"Hapur", 245101}, {"Meertt", 245304}},
		IsLoading: false,
		IsActive: true,
	}

	er := t.Execute(os.Stdout, p1)
	if(err != nil) {
		panic(er)
	}
}