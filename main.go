// Program By Adam Bouaazzi
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Person struct {
	Name          string `json:"name"`
	Age           int64  `json:"age"`
	Favoritecolor string `json:"favoritecolor"`
}

var p Person

// Create a map which will hold the Person Struct
// The map will create Keys and Values to match what attributes are in the struct
var m = make(map[string]Person)

// Users should be able to access all the people by querying with the people path
// etc http://localhost:8080/people/
// expected output would be
// [{"id":"1","Name":"Adam","favorite_color":"Blue","Age":22}]
// [{"id":"2","Name":"Scott","favorite_color":"Blue","Age":22}]
//func getPeople(w http.ResponseWriter, r *http.Request) {
//	var jsonData []byte
//	jsonData, err := json.Marshal(people)
//	if err != nil {
//		log.Println(err)
//	}
//	fmt.Println(string(jsonData))
//}

// Users should be able to access 1 person by making a request to /people/{name}
// etc http://localhost:8080/people/Adam
// expected output would be
//
// {
//     "name": "Adam",
//     "favoritecolor": "Blue",
//     "age": 22
// }
func getPeople(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := json.Marshal(m)
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	_, err = w.Write([]byte("GET ALL PEOPLE\n"))
	w.Write(jsonBody)

}

func getPerson(w http.ResponseWriter, r *http.Request) {
	//Update the URL to the current path
	fmt.Printf("url: %+v \n", r.URL)
	jsonData, err := json.Marshal(m)

	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	// Write map and info message to Postman
	_, err = w.Write([]byte("GET 1 PERSON\n"))
	w.Write(jsonData)
	if err != nil {
		fmt.Println("error writing the body")
	}

}
func createPerson(w http.ResponseWriter, r *http.Request) {

	// read the body of the request
	data, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(data))
	if err != nil {
		fmt.Println("err reached")
		w.WriteHeader(500)
	}

	//unmarshal from json to struct
	var people Person

	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Println("error during unmarshal", err)
		w.WriteHeader(500)
		w.Write(nil)
	}
	// Allow the key of the map to be set to the name attribute in the person struct
	m[people.Name] = people

	w.WriteHeader(200)
	//Write to the body of the program
	_, err = w.Write([]byte("POST REQUEST WAS SUCCESFUL\n"))
	w.Write(data)

	if err != nil {
		fmt.Println("person wasnt received")
	}

}

// The HandleRequests function will be responsible for determining wheter
// a GET/POST request in made on the url http://localhost:8080/people
func routeRequests(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listening on 8080")
	switch r.Method {
	case "GET":

		if r.URL.Path != "/people" {

			getPerson(w, r)
		}
		fmt.Printf("current url location: %+v \n", r.URL)
		getPeople(w, r)
	case "POST":
		createPerson(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

func main() {
	// Create a Handler which will create the handle requests function
	// This function will be responsible for GET/POST HTTP requests
	// The api will be hosted on http://localhost:8080/people

	http.HandleFunc("/people", routeRequests)
	http.HandleFunc("/people/", getPerson)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
