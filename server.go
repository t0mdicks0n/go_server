package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", defaultPath)
	http.HandleFunc("/api/data", returnData)
	
	fmt.Println("Server is listening to port 1337")
	http.ListenAndServe(":1337", nil)
}

func defaultPath(w http.ResponseWriter, r *http.Request) {
	fmt.Println("A web user requested the / path")
	http.ServeFile(w, r, r.URL.Path[1:])
}

func returnData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("A web user requested the /data path \n")
	fmt.Fprintf(w, "Your request was succesfull, I will get you your data \n")
}