// This is where the server will start
package main
import(
	"net/http"
	"github.com/go-chi/chi/v5"

	"github.com/wthobika/BarkingMelodies/src"
)

//Where the server starts
func main(){
	router := chi.NewRouter()

	//ROUTES
	router.Get("/", src.IndexHandler)
	router.Post("/register", src.Register)
	router.Post("/login", src.Login)
	router.Post("/logout", src.Logout)

	http.ListenAndServe(":8080", router)
}
