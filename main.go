package main

import (
	"fmt"
	"net/http"
)


func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello There Awesome Go Program</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello From Contact</h1>")
}


func main() {
  http.HandleFunc("/", homeHandler);
  http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting Web Server on 3000...");
	http.ListenAndServe(":3000", nil)
}