package config

import (
	"html/template"
	"net/http"
	"path"
)

func SetTemplate(writer http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // ini untuk cek, jika path / url gak ditemukan
		http.NotFound(writer, r)
		return
	}
	var filePath = path.Join("view", "index.html") // folder tidak boleh huruf besar...!
	var tmpl, err = template.ParseFiles(filePath)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Website Golang",
		"name":  "saya",
	}

	err = tmpl.Execute(writer, data)
	if err != nil {
		InternalServerError(writer)
	}
}
