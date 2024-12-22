// This is where the server will start
package main
import(
	"net/http"
	"github.com/go-chi/chi/v5"

	"github.com/wthobika/BarkingMelodies/src"     //I have no idea why this import wont work but it wont so whatever. Maybe I'll fix it in the future.
)

//Where the server starts
func main(){
	router := chi.NewRouter()

	//ROUTES
	router.Get("/", src.indexHandler)

	
	http.ListenAndServe(":8080", router)
}



//I couldn't figure out how to import files so I will write the logic/backend here temporarily

//Serves basic index.html file
// func indexHandler(w http.ResponseWriter, r *http.Request){
// 	http.ServeFile(w, r, "../static/html/index.html")
// }