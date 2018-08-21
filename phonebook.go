package main

import (
	"github.com/gin/controllers/routes"
	"github.com/gin/types"
)

func main() {

	types.People = append(types.People, types.Person{ID: "1", Firstname: "James", Lastname: "Bond", Address: &types.Address{City: "City X", State: "State X"}})
	types.People = append(types.People, types.Person{ID: "2", Firstname: "Cool", Lastname: "Bond", Address: &types.Address{City: "City Z", State: "State Y"}})
	types.People = append(types.People, types.Person{ID: "3", Firstname: "Hot", Lastname: "Bond"})

	r := routes.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
