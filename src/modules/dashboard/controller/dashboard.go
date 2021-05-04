package modules

import (
	"html/template"
	"net/http"
	"path"
)

func IndexDashboard(writer http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // ini untuk cek, jika path / url gak ditemukan
		http.NotFound(writer, r)
		return
	}
	var tmpl, err = template.ParseFiles(path.Join("modules/dashboard/views", "index.html"), path.Join("view", "default-Layout.html"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"petualangan": "saya",
	}

	err = tmpl.Execute(writer, data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
