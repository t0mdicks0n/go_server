package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Chats struct {
	ID string `json:"id,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Username string `json:username,omitempty`
	Message string `json:"message,omitempty"`
	Group string `json:"group,omitempty"`
}

var chatCache []Chats

func main() {
	chatCache = append(chatCache, Chats{ID: "1", Timestamp: "2017-07-21", Username: "Tom", Message: "Hello all", Group: "lobby"})
	chatCache = append(chatCache, Chats{ID: "2", Timestamp: "2017-07-21", Username: "Ebba", Message: "Hey brother", Group: "lobby"})
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
	json.NewEncoder(w).Encode(chatCache)
}