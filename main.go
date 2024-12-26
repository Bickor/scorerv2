package main

import (
	"database/sql"
	"fmt"

	// "fmt"
	"html/template"
	// "log"
	"net/http"
	// "strconv"
	// "strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var tmpl *template.Template
var db *sql.DB

func init() {
	var err error
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		fmt.Println("Error found: " + err.Error())
	}
}

func initDB() {
	return
}

func main() {
	gRouter := mux.NewRouter()

	// initDB()
	// defer db.Close()

	gRouter.HandleFunc("/", Homepage)

	http.ListenAndServe(":4000", gRouter)
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		fmt.Println("Error found: " + err.Error())
	}
}
