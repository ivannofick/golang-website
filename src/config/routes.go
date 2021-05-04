package config

import (
	"fmt"
	"net/http"
)

func HandlerRoutes() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", HandlerIndex)
	http.HandleFunc("/product", HandlerProductDetail)
	handlerServe()
}

/**
* Ini funct untuk handler serve (go run main.go)
**/
func handlerServe() {
	fmt.Println("Server Has Been Activated") // untuk mencetak apakah server sudah aktif
	http.ListenAndServe(":9000", nil)        // untuk menuliskan port saat kita akan mengaktifkan file
}
