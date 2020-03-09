package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/atrastudhi/go-rest-api/controllers"
)

func main()  {
	var router = mux.NewRouter()

	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.GetUsers).Methods("GET")

	fmt.Println("listen on port 3000...")
	http.ListenAndServe(":3000", router)
}