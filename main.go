package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"strconv"
)

type item struct {
	ID          int `json:ID`
	Name        string `json:"Name"`
	Price       uint64 `json:"Price"`
}

type allItems []item

var items = allItems{}

var shoppingCart []int

func loadJSON() {
	// read file
    data, err := ioutil.ReadFile("./inventory.json")
    if err != nil {
      fmt.Print(err)
    }

    // unmarshall it
    err = json.Unmarshal(data, &items)
    if err != nil {
        fmt.Println("error:", err)
    }
}

func addItemToCart(w http.ResponseWriter, r *http.Request) {
	itemID,err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
      fmt.Print(err)
    }

    /*
    	obviously there would be logic to validate the id. I am just adding it to the shopping card.
    	Also, there is a mono shopping card, which would be in the database not an array. 
    */
    
	shoppingCart = append(shoppingCart,itemID)

	json.NewEncoder(w).Encode(shoppingCart)
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func getShoppingCart(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(shoppingCart)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to vaporware ecommerce!")
}

func main() {
	loadJSON()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/items", getAllItems).Methods("GET")
	router.HandleFunc("/cart", getShoppingCart).Methods("GET")
	router.HandleFunc("/cart/{id}", addItemToCart).Methods("POST")
	var port = "4000"
	log.Println("Starting Server http://localhost:"+port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}