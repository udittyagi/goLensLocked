package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func faqHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  fmt.Fprint(w, `<h1>FAQ Page</h1>
  <ul>
    <li>
      <b>Is there a free version?</b>
      Yes! We offer a free trial for 30 days on any paid plans.
    </li>
    <li>
      <b>What are your support hours?</b>
      We have support staff answering emails 24/7, though response
      times may be a bit slower on weekends.
    </li>
    <li>
      <b>How do I contact support?</b>
      Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
    </li>
  </ul>
  `)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch(r.URL.Path) {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
  case "/faq":
		faqHandler(w, r)
	default: 
    notFoundHandler(w, r)
	}
}

func MyRequestHandler(w http.ResponseWriter, r *http.Request) {
  // fetch the url parameter `"userID"` from the request of a matching
  // routing pattern. An example routing pattern could be: /users/{userID}
  userID := chi.URLParam(r, "userID")
  fmt.Fprint(w, fmt.Sprintf("<h1>User Key is %v</h1>", userID))
}


func main() {
  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Get("/", homeHandler)
  r.Get("/contact", contactHandler)
  r.Get("/faq", faqHandler)
  r.Get("/users/{userID}", MyRequestHandler)

  r.NotFound(notFoundHandler)

	fmt.Println("Starting Web Server on 3000...");
	http.ListenAndServe(":3333", r)
}