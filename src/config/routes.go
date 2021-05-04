package config

import (
	"fmt"
	dashboard "golang-web-system/modules/dashboard/controller"
	"net/http"
)

func HandlerRoutes() {
	http.HandleFunc("/", dashboard.IndexDashboard)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	handlerServe()
}

/**
* Ini funct untuk handler serve (go run main.go)
**/
func handlerServe() {
	fmt.Println("Server Has Been Activated") // untuk mencetak apakah server sudah aktif
	http.ListenAndServe(":9000", nil)        // untuk menuliskan port saat kita akan mengaktifkan file
}
