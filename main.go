package main

import (
	"fmt"
	"net/http"
)


func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello There Awesome Go Program</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello From Contact Handler</h1>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
  http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}


//We can use instance of Router in http.ListenAndServe as 2nd argument, since it implements the http.Handler interface 
// by implementing ServerHttp method
type Router struct {}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  switch(r.URL.Path) {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default: 
    notFoundHandler(w, r)
	}
}

func main() {
  var router Router
	fmt.Println("Starting Web Server on 3000...");
	http.ListenAndServe(":3000", router)
}