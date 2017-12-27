package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

// Global Variables
var (
	sessionObj    *mgo.Session
	collectionObj *mgo.Collection
)

func main() {
	// Connecting to the mongodb server
	serverName := "localhost"
	sessionObj, err := mgo.Dial(serverName)
	defer sessionObj.Close()
	if err != nil {
		panic(err)
	}
	// fmt.Println("Pausing for 6 seconds")
	// time.Sleep(5 * time.Second)

	// creating a collection object
	collectionObj = sessionObj.DB("sampleDB").C("phonebook")

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

// Display all from the people var
func GetPeople(c *gin.Context) {
	var result []Person
	err := collectionObj.Find(nil).All(&result)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(c.Writer).Encode(result)
}

// Display a single data using id
func GetPerson(c *gin.Context) {
	id := c.Param("id")
	var result []Person
	err := collectionObj.Find(bson.M{"id": id}).All(&result)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(c.Writer).Encode(result)
}

// create a new item
func CreatePerson(c *gin.Context) {
	id := c.Param("id")
	var person Person
	person.ID = id

	// Decoding the json data to person struct
	_ = json.NewDecoder(c.Request.Body).Decode(&person)
	err := collectionObj.Insert(person)
	if err != nil {
		glog.Fatalf("Add fail %v\n", err)
	}
	json.NewEncoder(c.Writer).Encode(person)
}

// update an existing item
func UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var person Person
	_ = json.NewDecoder(c.Request.Body).Decode(&person)

	colQuerier := bson.M{"id": id}
	change := bson.M{"$set": bson.M{
		"id":        person.ID,
		"firstname": person.Firstname,
		"lastname":  person.Lastname}}
	_, _ = collectionObj.UpdateAll(colQuerier, change)
	json.NewEncoder(c.Writer).Encode(person)
}

// Delete an item using id
func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	_, err := collectionObj.RemoveAll(bson.M{"id": id})
	if err == nil {
		c.String(200, "Deleted item with id: %v", id)
	}
}
