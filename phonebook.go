package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/people", GetPeople)
	r.GET("/people/:id", GetPerson)
	r.POST("/people/:id", CreatePerson)
	r.PUT("/people/:id", UpdatePerson)
	r.DELETE("/people/:id", DeletePerson)

	return r
}

func main() {

	people = append(people, Person{ID: "1", Firstname: "James", Lastname: "Bond", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Cool", Lastname: "Bond", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Hot", Lastname: "Bond"})

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

// Display all from the people var
func GetPeople(c *gin.Context) {
	json.NewEncoder(c.Writer).Encode(people)
}

// Display a single data using id
func GetPerson(c *gin.Context) {
	id := c.Param("id")
	for _, item := range people {
		if item.ID == id {
			json.NewEncoder(c.Writer).Encode(item)
			return
		}
	}
	json.NewEncoder(c.Writer).Encode(&Person{})

}

// create a new item
func CreatePerson(c *gin.Context) {
	id := c.Param("id")
	var person Person
	_ = json.NewDecoder(c.Request.Body).Decode(&person)
	person.ID = id
	people = append(people, person)
	json.NewEncoder(c.Writer).Encode(people)
}

// update an existing item
func UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var person Person
	_ = json.NewDecoder(c.Request.Body).Decode(&person)
	for i := 0; i < len(people); i++ {
		if id == people[i].ID {
			people[i].ID = person.ID
			people[i].Firstname = person.Firstname
			people[i].Lastname = person.Lastname
		}
	}
	json.NewEncoder(c.Writer).Encode(people)
}

// Delete an item using id
func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	for index, item := range people {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(c.Writer).Encode(people)
	}
}
