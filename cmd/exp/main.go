package main

import (
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age int
}

func main() {
	t ,err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	p1 := Person {
		"Udit", 28, 
	}
	er := t.Execute(os.Stdout, p1)
	if(err != nil) {
		panic(er)
	}
}