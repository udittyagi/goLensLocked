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
  http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
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
  /*
    We can use pathHandler in ListenAndServe becoz pathHandler has same type format as http.HandlerFunc
    http.HandleFunc('/', pathHandler) method accepted the pathHandler as it accept function of  func(ResponseWriter, *Request) type.
    But in order to send this to ListenAndServe we need to type convert the pathHandler to http.HandlerFunc which is basically type
    with underlying type of  func(ResponseWriter, *Request), which our pathHandler matches and http.HandlerFunc also implements ServerHTTP method
    so http.HandlerFunc implements http.Handler interface as well hence by type conveting our pathHandler, it will be a good candidate
    to be passed in ListenAndServe
  */
  var router http.HandlerFunc = pathHandler
	fmt.Println("Starting Web Server on 3000...");
	http.ListenAndServe(":3000", router)
}