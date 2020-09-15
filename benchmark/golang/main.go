package main

import (
	"fmt"
	"log"
	"net/http"

	//"examples.com/ramlcode/goraml"

	"github.com/gorilla/mux"
	//"github.com/watercraft/validator"
)

func main() {
	// input validator
	//validator.SetValidationFunc("multipleOf", goraml.MultipleOf)

	r := mux.NewRouter()

	// home page
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w, r, "index.html")
		fmt.Fprintf(w, "Hello World!")
	})

	// apidocs
	r.PathPrefix("/apidocs/").Handler(http.StripPrefix("/apidocs/", http.FileServer(http.Dir("./apidocs/"))))

	UsersInterfaceRoutes(r, UsersAPI{})

	log.Println("starting server")
	http.ListenAndServe(":5000", r)
}
