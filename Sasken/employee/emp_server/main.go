package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

//Employee Model
type Employee struct {
	EmpID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EmpName string             `json:"empname,omitempty" bson:"empname",omitempty`
	MgrID   string             `json:"mgrid,omitempty" bson:"mgrid,omitempty"`
}

//Function to check error and print user defing msg alsong with error
func checkErr(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

// function to set the gorilla mux router and create the http server
func setupRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/employee", CreateEmployeeEndPoint).Methods("POST")
	router.HandleFunc("/employee/{id}", GetEmployeeEndpoint).Methods("GET")
	log.Println("listening to the port :8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}

//creatin of set of employees
func CreateEmployeeEndPoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var emp Employee
	_ = json.NewDecoder(request.Body).Decode(&emp)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, emp)
	fmt.Println("data Inserted Successfully ", result)
	json.NewEncoder(response).Encode(result)
}

//getting an employee detail and showing detail whom he resports
func GetEmployeeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	fmt.Println(params)
	oid, _ := primitive.ObjectIDFromHex(params["id"])
	emp := &Employee{}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.M{"_id": oid}
	err := collection.FindOne(ctx, filter).Decode(emp)
	fmt.Println(oid)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	fmt.Println("received the employee details successfully ", emp)
	json.NewEncoder(response).Encode(emp)
	//getting the manager name of a given employee
	mgrid, _ := primitive.ObjectIDFromHex(emp.MgrID)
	filter = bson.M{"_id": mgrid}
	err = collection.FindOne(ctx, filter).Decode(emp)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	fmt.Println("received the employee's manager details successfully ", emp)
	json.NewEncoder(response).Encode(emp)
}

func createDB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	collection = client.Database("mydb").Collection("employee")
}

// The application starts from here
func main() {
	// to print the file name and line number if code crash
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Starting the application...")
	//connection to the mongo DB and creating the employee collection/table
	createDB()
	setupRoutes()
	fmt.Println("disconnecting the db")
	defer client.Disconnect(context.TODO())
}
