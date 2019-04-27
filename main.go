package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const userDB = "root"
const passDB = "prezes11"
const nameDB = "game"

type Action struct {
	ID         int    `json:"id"`
	UserID     int    `json:"userId"`
	GameID     int    `json:"gameId"`
	ActionMove string `json:"actionMove"`
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func ShowAllAction(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", userDB+":"+passDB+"@/"+nameDB+"?parseTime=true")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM action")
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println(results)

	var actions []Action
	for results.Next() {
		var action Action

		err = results.Scan(&action.ID, &action.UserID, &action.GameID, &action.ActionMove)
		if err != nil {
			panic(err.Error())
		}
		actions = append(actions, Action{ID: action.ID, UserID: action.UserID, GameID: action.GameID, ActionMove: action.ActionMove})
		// log.Printf(action.ActionMove)
	}

	fmt.Println("Endpoint Hit: ShowAllAction")

	json.NewEncoder(w).Encode(actions)
}

func CreateAction(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	db, err := sql.Open("mysql", userDB+":"+passDB+"@/"+nameDB+"?parseTime=true")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO action (userId, gameId, action) VALUES (?, ?, ?)", r.Form["userId"][0], r.Form["gameId"][0], r.Form["action"][0])

	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer insert.Close()
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomePage)
	router.HandleFunc("/all", ShowAllAction).Methods("GET")
	router.HandleFunc("/create", CreateAction).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	handleRequests()
}
