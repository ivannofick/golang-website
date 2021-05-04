package config

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func MainRun() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", handlerIndex)

	fmt.Println("Server Has Been Activated") // untuk mencetak apakah server sudah aktif
	http.ListenAndServe(":9000", nil)        // untuk menuliskan port saat kita akan mengaktifkan file
}

func handlerIndex(writer http.ResponseWriter, r *http.Request) {
	var filePath = path.Join("View", "index.html")

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
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}