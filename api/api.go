package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gsm23/go-rest-api/common"
//	"github.com/gsm23/go-rest-api/config"
)


type App struct {
	Router *mux.Router
}

type (
	s = common.WebServer
)

func TestApi(a int) {
	fmt.Println("\nInside API Module")
	fmt.Println(a)

}

//func (a *App) MyApp(s *Server) {
//	fmt.Println("Starting the web server")
//	fmt.Println(&s)
	//a.Router = mux.NewRouter()
	//r.HandleFunc("/", showHomepage).Methods("GET")
	//log.Fatal(http.ListenAndServe(":8080", r))
	//a.initializeRouter()
	//fmt.Println(s.ContentType)
//}

/*
func showHomepage(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello From my API....."))
}

// Starting server
func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

func (a *App) initializeRouter() {
	//a.Router.HandleFunction("/", showHomepage).Method("GET")
	a.Router.Get("/").HandlerFunc(showHomepage)
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

 */