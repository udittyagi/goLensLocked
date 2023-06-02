package main

import (
	"html/template"
	"os"
)

type Person struct {
	Name string
	Bio string
	HtmlStr template.HTML
}

func main() {
	t ,err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	/**
		If we use html/template then Bio will get excaped  &lt;script&gt;alert(&#34;Hello There!!&#34;)&lt;/script&gt;%
		If we use text/templete then bio will not get escaped <script>alert("Hello There!!")</script>%

		What if we want to set the HTML in our template.HTML type,  like we are  using in struct HtmlStr field
	**/
	p1 := Person {
		Name: "Udit", 
		Bio: `<script>alert("Hello There!!")</script>`,
		HtmlStr: `<script>alert("Hello There!!")</script>`,
	}
	er := t.Execute(os.Stdout, p1)
	if(err != nil) {
		panic(er)
	}
}