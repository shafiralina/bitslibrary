package main

import (
	"bitslibrary/app"
	"bitslibrary/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/list", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/user/update/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/book/new", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/book/list", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/book/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/newest/book", controllers.GetNewestBooks).Methods("GET")
	router.HandleFunc("/api/book/update/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/stock/new", controllers.CreateStock).Methods("POST")
	router.HandleFunc("/api/borrow/new", controllers.CreateBorrow).Methods("POST")
	router.HandleFunc("/api/borrow/{id}", controllers.GetBorrow).Methods("GET")
	router.HandleFunc("/api/borrow/detail/{id}", controllers.GetDetailBorrow).Methods("GET")

	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Started at: " + port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
