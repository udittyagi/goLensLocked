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


func homeHandler(w http.ResponseWriter, r *http.Request) {

  // Same as node.js path module
	 filePath := filepath.Join("templates", "home.gohtml")
   tpl, err := template.ParseFiles(filePath);

   //This error occurs when our template fails to parse
   //This can happen for a variety of reasons, but they all relate to a template that isn’t valid no matter what data we provide to it
   if err != nil {
    log.Printf("Error Occured While parsing template: %v", err)
    http.Error(w, "Error Occured while parsing template", http.StatusInternalServerError);
    return;
   }

   erEx := tpl.Execute(w, nil);

   //occur when we attempt to execute our template
   // This type of error can occur if our template parses correctly, but execution fails for some reason
   // eg: if the data we pass in is missing a field our template needs
   if erEx != nil {
    log.Printf("Error Occured While executing template: %v", err);

    //Event though we are sending error here, there might be the case
    //where tpl.Execute(w, nil); have written some data on request stream previously
    // so we may see that data too on client
    http.Error(w, "Error Occured while executing template", http.StatusInternalServerError);
   }
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