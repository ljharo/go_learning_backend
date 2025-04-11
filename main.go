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
	fmt.Fprintf(w, "Post Users")
}

func PutUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Put Users")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Delete Users")

}


func main (){
	r := mux.NewRouter().StrictSlash(false)

	// user CRUD
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUsers).Methods("GET")
	r.HandleFunc("/users", PostUsers).Methods("POST")
	r.HandleFunc("/users/{id}", PutUsers).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUsers).Methods("DELETE")

	server := &http.Server{
		Addr: ":8000",
		Handler: r,
		ReadTimeout:10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	log.Fatal(server.ListenAndServe())
}