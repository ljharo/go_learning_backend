package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func animation(retraso time.Duration){
	for {
		for _, r := range `\|/-`{
			fmt.Printf("\r%c", r)
			time.Sleep(retraso)
		}
	}
}

func home (w http.ResponseWriter, r *http.Request){
	fmt.Println("Request received from home")
	fmt.Fprintf(w, "<h1>hola mundo</h1>")
}

func test (w http.ResponseWriter, r *http.Request){
	fmt.Println("Request received from test")
	fmt.Fprintf(w, "<h1>hola mundo from test</h1>")
}

type mensaje struct {
	msg string

}

func (m mensaje) ServeHTTP (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, m.msg)
}

func main (){

	mux := http.NewServeMux()
	var msg mensaje = mensaje{"hola mundo"}
	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/", fs)
	mux.HandleFunc("/test", test)
	mux.Handle("/hola", msg)


	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening ...")
	log.Fatal(server.ListenAndServe())

}