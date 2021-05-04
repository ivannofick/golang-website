package config

import (
	"fmt"
	"net/http"
	"strconv"
)

func HandlerProduct(writer http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNum, err := strconv.Atoi(id)
	if err != nil || idNum < 1 {
		http.NotFound(writer, r)
		return
	}

	fmt.Fprintf(writer, "Product Page : %d", idNum)

	// writer.Write([]byte("Product Page"))

}
