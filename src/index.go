package src

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
	//The file path is with relation to the main.go file, not the file in src. Thats why we use static/ instead of ../static/
	http.ServeFile(w, r, "static/html/index.html")
	fmt.Println("Index served.")
}