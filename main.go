package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type form struct {
	Mode string
	Pw   string
	Id   string
}

func main() {
	result := template.Must(template.ParseFiles("template/result.html"))
	test := template.Must(template.ParseFiles("template/test.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		data := form{Mode: r.Method}
		if r.Method == "GET" {
			data.Id = r.URL.Query().Get("id")
			data.Pw = r.URL.Query().Get("pw")
		} else if r.Method == "POST" {
			data.Id = r.FormValue("id")
			data.Pw = r.FormValue("pw")
		}
		result.Execute(w, data)
	})
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		result := ""
		for key, value := range r.URL.Query() {
			result += fmt.Sprintf("%s : %s<br>", key, value[0])
		}
		data := struct {
			Result string
		}{
			Result: result,
		}
		test.Execute(w, data)
	})
	http.ListenAndServe(":1234", nil)
}
