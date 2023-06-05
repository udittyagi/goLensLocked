package main

import (
	"html/template"
	"os"
)

type Person struct {
	Name string
	Bio string
	Age int
}

func main() {
	t ,err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	/**
		Go encode the data based on where it present, if it is in html it is encoded differently
		if it is in script tag, it is encoded differently
		In template Bio will be  &lt;script&gt;alert(&#34;Hello There!!&#34;)&lt;/script&gt;%
		Inside script Bio will be "\u003cscript\u003ealert(\"Hello There!!\")\u003c/script\u003e"
		Inside script the data type is also taken into consideration, Age field will be int in script
		whereas string in template
	**/
	p1 := Person {
		Name: "Udit", 
		Bio: `<script>alert("Hello There!!")</script>`,
		Age: 24,
	}
	er := t.Execute(os.Stdout, p1)
	if(err != nil) {
		panic(er)
	}
}