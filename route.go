package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var vue int = 0

func main() {
	temp, err := template.ParseGlob("./temp/*.html")
	if err != nil {
		fmt.Println(fmt.Sprintf("Erreur %s", err.Error()))
	}
	type PageVariables struct {
		Nom     string
		Filiere string
		Niv     int
		Nbr     int
	}

	type Eleve struct {
		Prenom string
		Nom    string
		Age    int
		Sexe   bool
	}

	type data struct {
		PV  PageVariables
		Elv []Eleve
	}

	type even struct {
		Value int
		Check bool
	}

	type person struct {
		Nom string
		Prenom string
		Birthday string
		Sexe string
	}

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		vue++
		var page even
		if vue%2 == 0 {
			page = even{vue, true}
		} else {
			page = even{vue, false}
		}
		temp.ExecuteTemplate(w, "change", page)
	})

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		lstelv := []Eleve{{"Cyril", "RODRIGUES", 22, true}, {"Kheir-eddine", "MEDERREG", 22, false}, {"Alan", "PHILIPIERT", 26, true}}
		page := PageVariables{"Mentor'ac", "Informatique", 5, len(lstelv)}
		d := data{page, lstelv}
		temp.ExecuteTemplate(w, "promo", d)
	})

	http.HandleFunc("/user/init", func(w http.ResponseWriter, r *http.Request) {

	})
	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {

	})
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8080", nil)
}
