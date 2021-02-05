package api

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}


func (a *App) MyApp() {
	fmt.Println("Starting the web server")
	a.Router = mux.NewRouter()
	//r.HandleFunc("/", showHomepage).Methods("GET")
	//log.Fatal(http.ListenAndServe(":8080", r))
	a.initializeRouter()
}


func showHomepage(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello From my API....."))
}

// Starting server
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

func (a *App) initializeRouter() {
	a.Router.HandleFunction("/", showHomepage)
}

//Error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

//Response writer as JSON format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}