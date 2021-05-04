package modules

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func IndexProfile(writer http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNum, err := strconv.Atoi(id)
	if err != nil || idNum < 1 {
		http.NotFound(writer, r)
	}

	data := map[string]interface{}{
		"content": idNum,
	}
	// var tmpl, err = template.ParseFiles(path.Join("modules/dashboard/views", "index.html"), path.Join("view", "default-Layout.html"))

	tmpl, err := template.ParseFiles(path.Join("modules/profile/views", "index.html"), path.Join("../../src/view", "defaul-layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(writer, "Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(writer, data)
	if err != nil {
		log.Println(err)
		http.Error(writer, "Error", http.StatusInternalServerError)
		return
	}
}
