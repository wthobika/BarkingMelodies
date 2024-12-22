// This is where the server will start
package main
import(
	"net/http"
	"github.com/go-chi/chi/v5"
)


func main(){
	router := chi.NewRouter()

	//ROUTES
	router.Get("/", indexHandler)

	
	http.ListenAndServe(":8080", router)
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "../static/html/index.html")
}