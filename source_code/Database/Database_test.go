package Database

import (
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Dob         time.Time `json:"dob"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	CreateAt    time.Time
}

// Insertion test
func TestInsert(t *testing.T) {
	fmt.Println("\nINSERTION TEST")
	var opt User
	ipt := User{"123456", "XYZ", time.Date(2000, 10, 01, 0, 0, 0, 0, time.UTC), "Mysore", "Student", time.Now()}
	client, _ := DBconnect()
	err := InsertData(client, &ipt)
	if err != nil {
		t.Error("Error should not have occurred")
	}
	RetriveData(client, "123456", &opt)
	fmt.Println("Input given : ", ipt)
	fmt.Println("Output obtained :", opt)
	DBdisconnect(client)
	fmt.Println()
}

// Deletion test
func TestDeletion(t *testing.T) {
	fmt.Println("\nDELETION TEST")
	var opt User
	id := "123456"
	client, _ := DBconnect()
	RetriveData(client, "123456", &opt)
	err := DeleteDBData(client, id)
	if err != nil {
		t.Error("Error should not have occurred")
	}
	fmt.Println("Deleted record : ", opt)
	DBdisconnect(client)
	fmt.Println()
}

// Retrival Test
func TestRetrive(t *testing.T) {
	fmt.Println("\nRETRIVAL TEST")
	var opt User
	ipt := User{"123456", "XYZ", time.Date(2000, 10, 01, 0, 0, 0, 0, time.UTC), "Mysore", "Student", time.Now()}
	client, _ := DBconnect()
	InsertData(client, &ipt)
	err := RetriveData(client, "123456", &opt)
	if err != nil {
		t.Error("Error should not have occurred")
	}
	fmt.Println("Input given : ", ipt)
	fmt.Println("Output obtained :", opt)
	DBdisconnect(client)
	fmt.Println()
}

// Updation Test
func TestUpdate(t *testing.T) {
	fmt.Println("\nUPDATION TEST")
	var opt User
	ipt := User{"123456", "XYZ", time.Date(2000, 10, 01, 0, 0, 0, 0, time.UTC), "Mysore", "Student", time.Now()}
	client, _ := DBconnect()
	RetriveData(client, "123456", &opt)
	filter := bson.D{{"id", "123456"}}
	update := bson.D{{"$set", bson.D{{"name", "XYZ"}, {"dob", time.Date(2000, 10, 01, 0, 0, 0, 0, time.UTC)}, {"address", "Bangalore"}, {"description", "Graduate"}, {"createAt", time.Now()}}}}
	err := UpdateDBData(client, filter, update)
	if err != nil {
		t.Error("Error should not have occurred")
	}
	RetriveData(client, "123456", &ipt)
	fmt.Println("previous output : ", opt)
	fmt.Println("updated output :", ipt)
	DBdisconnect(client)
	fmt.Println()
}
