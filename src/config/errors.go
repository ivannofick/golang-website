package config

import "net/http"

func InternalServerError(write http.ResponseWriter) {
	write.Write([]byte("Internal server error"))
}
