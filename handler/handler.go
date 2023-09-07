package handler

import (
	"app/api"
	"app/db"
	"html/template"
	"log"
	"net/http"
)

func IndexHandler(db db.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		todos, err := db.AllTodos()
		if err != nil {
			log.Println(err)
		}
		tmpl.Execute(w, map[string][]api.Todo{"Todos": todos})
	}
}

func AddTodoHandler(db db.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		description := r.PostFormValue("description")
		tmpl := template.Must(template.ParseFiles("index.html"))

		todo := api.Todo{Title: title, Description: description}
		db.AddTodo(todo)
		tmpl.ExecuteTemplate(w, "todo-list-item", todo)
	}
}
