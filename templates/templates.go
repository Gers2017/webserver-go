package templates

import (
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle  string
	Todos      []Todo
	PageFooter string
}

func RunTemplate() {
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
		PageFooter: "Made with golang",
	}

	tmpl := template.Must(template.ParseFiles("public/layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/todo", func(res http.ResponseWriter, req *http.Request) {
		q := req.URL.Query()
		title := q.Get("title")
		done := q.Get("done")
		isDone := done == "true" || done == "1"

		for index, todo := range data.Todos {
			if todo.Title == title {
				data.Todos[index].Done = isDone
			}
		}
		res.WriteHeader(200)
		fmt.Fprintf(res, "%s = %t", title, isDone)
	})
	port := ":80"
	fmt.Printf("Serving files at http://localhost%s/static/\n", port)
	http.ListenAndServe(port, nil)
}
