package codinggoweb

import (
	"html/template"
	"net/http"
)

// Todo : struct for template
type Todo struct {
	Title string
	Done  bool
}

// TodoPageData : struct to have all todos collectively
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

// BasicTemplateHTTP : function to use basic templating features
func BasicTemplateHTTP() {
	tmpl := template.Must(template.ParseFiles("./templates/layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO List",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
