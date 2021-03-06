package configs

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"testAPI/utils/environtment"
)

func CreateRouter() *mux.Router  {
	router := mux.NewRouter()
	return router
}
func RunServer(router *mux.Router)  {
	routerHost := environtment.Get("RouterHost", "localhost")
	routerPort := environtment.Get("RouterPort", "8080")

	log.Printf("Server is now listening at %v......\n", routerPort)
	err := http.ListenAndServe(routerHost+":"+routerPort, router)
	if err != nil{
		log.Print(err)
		log.Fatal(err)
	}
}
