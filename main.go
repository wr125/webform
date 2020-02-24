package main

import ("fmt"
"log"
"net/http"
"html/template"
"strconv"
)

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseGlob("Template/*.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil{
      log.Fatal(err)
    }
    for key, value := range r.Form{
      fmt.Printf("%s\n %v\n", key, value)
    }

    tpl.ExecuteTemplate(w, "index.html", nil)
}



func results(w http.ResponseWriter, r *http.Request) {

    if r.Method == "Post" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return

}
speed, err := strconv.ParseFloat(r.FormValue("speed")[0:], 10)
if err != nil {
  log.Fatal(err)
}
distance, err := strconv.ParseFloat(r.FormValue("distance")[0:], 10)
if err != nil {
  log.Fatal(err)
}
//example of wanting to show results on the html/template
fmt.Printf("%v\n", speed / distance)


car := struct{
  Speed float64
  Distance float64
}{
  Speed: speed,
  Distance: distance,
}

tpl.ExecuteTemplate(w, "results.html", car)


}
func main() {

  mux := http.NewServeMux()
      mux.HandleFunc("/", index)
      mux.HandleFunc("/results", results)

      fmt.Println("server is running..")
      err := http.ListenAndServe(":4000", mux)
      log.Fatal(err)
}
