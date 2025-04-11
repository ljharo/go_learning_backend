package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Get Users")
}

func PostUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Get Users")
}

func PutUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Get Users")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Get Users")

}


func main (){
	r := mux.NewRouter().StrictSlash(false)

	// user CRUD
	r.HandleFunc("api/user", GetUsers).Methods("GET")
	r.HandleFunc("api/user", PostUsers).Methods("POST")
	r.HandleFunc("api/user", PutUsers).Methods("PUT")
	r.HandleFunc("api/user", DeleteUsers).Methods("DELETE")

	server := &http.Server{
		Addr: "8000",
		Handler: r,
		ReadTimeout:10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	server.ListenAndServe()

}