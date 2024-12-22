package src

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "../static/html/index.html")
}