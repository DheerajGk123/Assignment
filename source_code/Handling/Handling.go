package Handling

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Assignment/source_code/Database"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// Schema
type User struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Dob         time.Time `json:"dob"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	CreateAt    time.Time
}

// Handle Creating Data
func CreateData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing CreateData()")
	var u1 User                          // Create a user object
	reqBody, _ := ioutil.ReadAll(r.Body) // Read POST parameters
	json.Unmarshal(reqBody, &u1)         // save the data in "reqBody" to the instance "u1"
	u1.CreateAt = time.Now()
	fmt.Println("u1 :", u1)

	// Open Database Connection
	client, err := Database.DBconnect()
	if err != nil {
		fmt.Println(err)
	}

	// Insert data into database
	err = Database.InsertData(client, &u1)
	if err != nil {
		fmt.Println(err)
	}

	// Close Database Connection
	err = Database.DBdisconnect(client)
	if err != nil {
		fmt.Println(err)
	}
}

// Handle Reading Data
func ReadData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing ReadData()")
	params := mux.Vars(r) // read URL parameters
	id := params["ID"]    // Get the ID
	fmt.Println("The ID is :", id)

	// Open Database Connection
	client, err := Database.DBconnect()
	if err != nil {
		fmt.Println(err)
	}

	// Retrive data from database
	var Robj User
	err = Database.RetriveData(client, id, &Robj)
	if err != nil {
		fmt.Println(err)
	}

	// Close Database Connection
	err = Database.DBdisconnect(client)
	if err != nil {
		fmt.Println(err)
	}

	// Send the fetched data
	fmt.Println("Sending data back to client")
	fmt.Println("Robj is :", Robj)
	json.NewEncoder(w).Encode(Robj)
}

// Handle Updating Data
func UpdateData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing UpdateData()")
	params := mux.Vars(r) // read URL parameters
	id := params["ID"]    // Get the ID
	fmt.Println("The ID is :", id)

	// Decode Update Data and store it in a structure
	var Uobj User
	json.NewDecoder(r.Body).Decode(&Uobj)
	filter := bson.D{{"id", Uobj.Id}}
	update := bson.D{{"$set", bson.D{{"name", Uobj.Name}, {"dob", Uobj.Dob}, {"address", Uobj.Address}, {"description", Uobj.Description}, {"createAt", Uobj.CreateAt}}}}

	// Open Database Connection
	client, err := Database.DBconnect()
	if err != nil {
		fmt.Println(err)
	}

	// Update data from database
	err = Database.UpdateDBData(client, filter, update)
	if err != nil {
		fmt.Println(err)
	}

	// Close Database Connection
	err = Database.DBdisconnect(client)
	if err != nil {
		fmt.Println(err)
	}
}

// Handle Deleting Data
func DeleteData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing DeleteData()")
	params := mux.Vars(r) // read URL parameters
	id := params["ID"]    // Get the ID
	fmt.Println("The ID is :", id)

	// Open Database Connection
	client, err := Database.DBconnect()
	if err != nil {
		fmt.Println(err)
	}

	// Delete data from database
	err = Database.DeleteDBData(client, id)
	if err != nil {
		fmt.Println(err)
	}

	// Close Database Connection
	err = Database.DBdisconnect(client)
	if err != nil {
		fmt.Println(err)
	}
}
