package server

import (
	"log"
	"net/http"

	"github.com/fidelicash/fc-api/logs"
	"github.com/fidelicash/fc-api/user"
	"github.com/fidelicash/fc-api/util"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var name string

func handlerOK(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola, Seja bem vindo ao fidelicash @" + name + " !!"))
}

// Listen init a http server
func Listen() {
	port := util.GetOSEnvironment("PORT", "5001")

	name = util.GetOSEnvironment("NAME", "JC")

	r := mux.NewRouter()
	r.Use(logs.LoggingMiddleware)

	user.SetRoutes(r.PathPrefix("/users").Subrouter())

	r.HandleFunc("/", handlerOK)
	http.Handle("/", r)

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Client-ID", "Content-Type", "X-Requested-With"})
	methodsOk := handlers.AllowedMethods([]string{"OPTIONS", "DELETE", "GET", "HEAD", "POST", "PUT"})

	log.Println("Listen on port: " + port)
	if err := http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(r)); err != nil {
		log.Fatal(err)
	}

	// log.Println("Listen on port: " + port)
	// if err := http.ListenAndServe(":"+port, r); err != nil {
	// 	log.Fatal(err)
	// }
}
