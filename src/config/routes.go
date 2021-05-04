package config

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func HandlerRoutes() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", HandlerIndex)
	http.HandleFunc("/product", HandlerProduct)
	fmt.Println("Server Has Been Activated") // untuk mencetak apakah server sudah aktif
	http.ListenAndServe(":9000", nil)        // untuk menuliskan port saat kita akan mengaktifkan file
}

func HandlerIndex(writer http.ResponseWriter, r *http.Request) {
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

func HandlerAbout(writer http.ResponseWriter, r *http.Request) {
	writer.Write([]byte("About Page"))
}
