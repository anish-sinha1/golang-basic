package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//JSON Marshalling

type person struct {
	First string
}

func main() {

	// p2 := person{
	// 	First: "James",
	// }
	// xp := []person{p1, p2}

	// bs, err := json.Marshal(xp)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println("Marshalled JSON", string(bs))

	// xp2 := []person{}
	// err = json.Unmarshal(bs, &xp2)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println("Back into struct", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Jenny",
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoding error")
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)

	if err != nil {
		log.Println("Unable to decode json")
	}

	log.Println("Decoded: ", p1)
}
