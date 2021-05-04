package config

import "net/http"

func HandlerProduct(writer http.ResponseWriter, r *http.Request) {
	writer.Write([]byte("Product Page"))

}
