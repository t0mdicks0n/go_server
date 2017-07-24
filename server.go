package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"

	"database/sql"
	_ "github.com/lib/pq"
)

// Create our test-data, will be replaced with persistant DB later on
type Chats struct {
	ID string `json:"id,omitempty"`
	Username string `json:username,omitempty`
	Message string `json:"message,omitempty"`
	Group string `json:"group,omitempty"`
}

var chatCache []Chats

func main() {
	router := mux.NewRouter()
	chatCache = append(chatCache, Chats{ID: "1", Username: "Tom", Message: "Hello all", Group: "lobby"})
	chatCache = append(chatCache, Chats{ID: "2", Username: "Ebba", Message: "Hey brother", Group: "lobby"})
	// Routes
	router.HandleFunc("/", defaultPath).Methods("GET")
	router.HandleFunc("/api/data", returnAllData).Methods("GET")
	router.HandleFunc("/api/data", createChatMsg).Methods("POST")
	router.HandleFunc("/api/data/{room}", getChatsForRooms).Methods("GET")
	router.HandleFunc("/api/data/{id}", deleteAChat).Methods("DELETE")
	// Initalize the server
	fmt.Println("Server is listening to port 1337")
	http.ListenAndServe(":1337", router)
}

func defaultPath(w http.ResponseWriter, r *http.Request) {
	fmt.Println("A web user requested the / path")
	var welcomeString string = "Welcome, try some of the API endpoints."
	json.NewEncoder(w).Encode(welcomeString)
}

func returnAllData(w http.ResponseWriter, r *http.Request) {
	// Access database
	db, err := sql.Open("postgres", "user=postgres dbname=chat sslmode=disable")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT id, username, message, groupname FROM chats;")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var allChats []Chats

	for rows.Next() {
		var (
			id string
			username string
			message string
			groupname string
		)

		rows.Scan(&id, &username, &message, &groupname)
		allChats = append(allChats, Chats{ID: id, Username: username, Message: message, Group: groupname})
	}
  db.Close()
  // Return the data to the request
	json.NewEncoder(w).Encode(allChats)
}

func createChatMsg (w http.ResponseWriter, r *http.Request) {
	var chat Chats
	json.NewDecoder(r.Body).Decode(&chat)
	fmt.Println("Writing chat message to DB from user: ", chat.Username)

	db, err := sql.Open("postgres", "user=postgres dbname=chat sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.QueryRow("INSERT INTO chats(username, message, groupname) VALUES ('" + chat.Username + "', ' " + chat.Message + "', '" + chat.Group + "');")

	db.Close()
	json.NewEncoder(w).Encode(chat)
}

func getChatsForRooms(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var chatsForRoom []Chats
	for _, chat := range chatCache {
		if chat.Group == params["room"] {
			chatsForRoom = append(chatsForRoom, chat)
		}
	}
	json.NewEncoder(w).Encode(chatsForRoom)
	return
}

func deleteAChat (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, chat := range chatCache {
		if chat.ID == params["id"] {
			chatCache = append(chatCache[:index], chatCache[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(chatCache)
}
