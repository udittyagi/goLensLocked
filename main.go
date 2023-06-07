package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate (w http.ResponseWriter, filePath string) {
  tpl, err := template.ParseFiles(filePath);

  if err != nil {
   log.Printf("Error Occured While parsing template: %v", err)
   http.Error(w, "Error Occured while parsing template", http.StatusInternalServerError);
   return;
  }

  err = tpl.Execute(w, nil);

  if err != nil {
   log.Printf("Error Occured While executing template: %v", err);
   http.Error(w, "Error Occured while executing template", http.StatusInternalServerError);
  }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("templates", "home.gohtml")
  executeTemplate(w, filePath);
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
  filePath := filepath.Join("templates", "contact.gohtml")
  executeTemplate(w, filePath);
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
  executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
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