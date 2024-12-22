// This is where the server will start
package main
import(
	"net/http"
	"github.com/go-chi/chi/v5"

	"github.com/wthobika/BarkingMelodies/src"
)


func main(){
	router := chi.NewRouter()

	//ROUTES
	router.Get("/", src.indexHandler)

	
	http.ListenAndServe(":8080", router)
}

// func indexHandler(w http.ResponseWriter, r *http.Request){
// 	http.ServeFile(w, r, "../static/html/index.html")
// }