package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"gopkg.in/go-playground/validator.v9"

)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("Template/*.html"))
}

func index(w http.ResponseWriter, r *http.Request) {



	tpl.ExecuteTemplate(w, "index.html", nil)
}

type Car struct {
	Speed    float64 `validate:"required"`
	Distance float64 `validate:"required"`
}

func (c Car) Acceleration() float64 {
	return c.Speed / c.Distance
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
	var validate *validator.Validate

	car := Car{Speed: speed, Distance: distance}

	validate = validator.New()
		err = validate.Struct(car)
		if err != nil {
			http.Error(w, err.Error(), 500)
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
