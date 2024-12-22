package src

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "../static/html/index.html")
}