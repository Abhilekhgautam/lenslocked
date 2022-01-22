package main

import(
  "fmt"
  "net/http"
)

func handlerFunc(w http.ResponseWriter, r* http.Request){
   w.Header().Set("Content-Type", "text/html")
//   fmt.Fprint(w,r.URL.Path)
   if r.URL.Path == "/"{
       fmt.Fprint(w, "<h1>Welcome to my awesome site! </h1>")
   } else if r.URL.Path == "/contact"{
      fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mail.to:supportlenslocked.com\">support@lenslocked.com</a>") 
 } else{
    w.WriteHeader(http.StatusNotFound)
   fmt.Fprint(w, "<h1>404 Not found</h1> <div>we cannot find the page</div>")
 }
}
func main(){

  http.HandleFunc("/", handlerFunc)
  http.ListenAndServe(":3000", nil)

}
