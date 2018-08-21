package people

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gin/types"
)

// Display all from the people var
func GetPeople(c *gin.Context) {
	json.NewEncoder(c.Writer).Encode(types.People)
}

// Display a single data using id
func GetPerson(c *gin.Context) {
	id := c.Param("id")
	for _, item := range types.People {
		if item.ID == id {
			json.NewEncoder(c.Writer).Encode(item)
			return
		}
	}
	json.NewEncoder(c.Writer).Encode(&types.Person{})

}

// create a new item
func CreatePerson(c *gin.Context) {
	id := c.Param("id")
	var person types.Person
	_ = json.NewDecoder(c.Request.Body).Decode(&person)
	person.ID = id
	types.People = append(types.People, person)
	json.NewEncoder(c.Writer).Encode(types.People)
}

// update an existing item
func UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var person types.Person
	_ = json.NewDecoder(c.Request.Body).Decode(&person)
	for i := 0; i < len(types.People); i++ {
		if id == types.People[i].ID {
			types.People[i].ID = person.ID
			types.People[i].Firstname = person.Firstname
			types.People[i].Lastname = person.Lastname
		}
	}
	json.NewEncoder(c.Writer).Encode(types.People)
}

// Delete an item using id
func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	for index, item := range types.People {
		if item.ID == id {
			types.People = append(types.People[:index], types.People[index+1:]...)
			break
		}
		json.NewEncoder(c.Writer).Encode(types.People)
	}
}
