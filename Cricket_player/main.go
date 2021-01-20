package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Player struct {
	Id      int
	Name    string
	Country string
	Role    string
	Age     int
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "order_db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Player ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}
	pl := Player{}
	res := []Player{}
	for selDB.Next() {
		var id, age int
		var name, country, role string
		err = selDB.Scan(&id, &name, &country, &role, &age)
		if err != nil {
			panic(err.Error())
		}
		pl.Id = id
		pl.Name = name
		pl.Country = country
		pl.Role = role
		pl.Age = age
		res = append(res, pl)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Player WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	pl := Player{}
	for selDB.Next() {
		var id, age int
		var name, country, role string
		err = selDB.Scan(&id, &name, &country, &role, &age)
		if err != nil {
			panic(err.Error())
		}
		pl.Id = id
		pl.Name = name
		pl.Country = country
		pl.Role = role
		pl.Age = age

	}
	tmpl.ExecuteTemplate(w, "Show", pl)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Player WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	pl := Player{}
	for selDB.Next() {
		var id, age int
		var name, country, role string
		err = selDB.Scan(&id, &name, &country, &role, &age)
		if err != nil {
			panic(err.Error())
		}
		pl.Id = id
		pl.Name = name
		pl.Country = country
		pl.Role = role
		pl.Age = age
	}
	tmpl.ExecuteTemplate(w, "Edit", pl)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		country := r.FormValue("country")
		role := r.FormValue("role")
		age := r.FormValue("age")
		insForm, err := db.Prepare("INSERT INTO player(name, country, role, age) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		log.Println("INSERT: Name: " + name + " | Country: " + country + " | Role:" + role + " | Age:" + age)
		insForm.Exec(name, country, role, age)

	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		country := r.FormValue("country")
		role := r.FormValue("role")
		age := r.FormValue("age")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Player SET name=?, country=?, role=?, age=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		// insForm.Exec(name, country, role, age, id)
		log.Println("INSERT: Name: " + name + " | Country: " + country + " | Role:" + role + " | Age:" + age)
		insForm.Exec(name, country, role, age, id)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	pl := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Player WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(pl)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8080")

	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
