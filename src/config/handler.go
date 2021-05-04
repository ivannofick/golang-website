package config

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

func HandlerIndex(writer http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // ini untuk cek, jika path / url gak ditemukan
		http.NotFound(writer, r)
		return
	}
	var tmpl, err = template.ParseFiles(path.Join("view", "index.html"), path.Join("view", "default-Layout.html"))
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

func HandlerAbout(writer http.ResponseWriter, r *http.Request) {
	writer.Write([]byte("About Page"))
}

func HandlerProductDetail(writer http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNum, err := strconv.Atoi(id)
	if err != nil || idNum < 1 {
		http.NotFound(writer, r)
		return
	}

	fmt.Fprintf(writer, "Product Page : %d", idNum)
}

func HandlerProduct(writer http.ResponseWriter, r *http.Request) {
	var tmpl, err = template.ParseFiles(path.Join("view", "index.html"), path.Join("view", "default-Layout.html"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Product Golang",
		"name":  "saya",
	}

	err = tmpl.Execute(writer, data)
	if err != nil {
		InternalServerError(writer)
	}
}
