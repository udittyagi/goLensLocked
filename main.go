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

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
  // w.WriteHeader(http.StatusNotFound)
  // fmt.Fprint(w, "Page Not Found")

  //We can also use http.Error
  // http.Error(w, "Page Not Found 2", http.StatusNotFound)

  http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func pathHanlder(w http.ResponseWriter, r *http.Request) {
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
  //HandleFunc are used only when DefaultServeMux is used, which in turn is used when nil is passed as second argument
  // in  ==> http.ListenAndServe
	http.HandleFunc("/", pathHanlder)

	fmt.Println("Starting Web Server on 3000...");
	http.ListenAndServe(":3000", nil)
}