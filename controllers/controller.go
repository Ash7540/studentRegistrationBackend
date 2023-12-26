package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"studentregist/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://ashwin123:ashwin123@cluster0.lie6s2w.mongodb.net/?retryWrites=true&w=majority"
const dbName = "student"
const colName = "studentForm"

var collection *mongo.Collection

// connecting with MongoDB
func init() {
	// set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// check the connection
	fmt.Println("MongoDB connection successful")

	collection = client.Database(dbName).Collection(colName)
	//collection instance
	fmt.Println("Collection instance created")
}

// mongodb helper

// insert one data
func insertSingleData(student models.StudentForm) {
	insertResult, err := collection.InsertOne(context.Background(), student)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

// insert multiple data
func insertMultipleData(students []interface{}) {
	insertManyResult, err := collection.InsertMany(context.Background(), students)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

// update data
func updateSingleData(studentId string) {
	id, _ := primitive.ObjectIDFromHex(studentId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"field_to_update": "new_value"}} // Update fields as needed
	// Perform the update operation

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data updated successfully")
}

// delete one data
func deleteSingleData(studentId string) {
	id, _ := primitive.ObjectIDFromHex(studentId)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents\n", result.DeletedCount)
}

// delete all data
func deleteAllData() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted documents in the collection with IDs:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// find data by ID
func findDataById(studentId string) models.StudentForm {
	var student models.StudentForm
	id, _ := primitive.ObjectIDFromHex(studentId)
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.Background(), filter).Decode(&student)
	if err != nil {
		log.Fatal(err)
	}
	return student
}

// get all data
func getAllData() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var students []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, movie)
	}

	defer cur.Close(context.Background())
	return students

}

// controller part

// get all student data
func GetAllStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allStudent := getAllData()
	json.NewEncoder(w).Encode(allStudent)
}

// get a single student data
func GetSingleStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	studentId := r.URL.Query().Get("id")
	student := findDataById(studentId)
	json.NewEncoder(w).Encode(student)
}

// create student data
func CreateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var student models.StudentForm
	_ = json.NewDecoder(r.Body).Decode(&student)
	insertSingleData(student)
	json.NewEncoder(w).Encode(student)
}

// create multiple student data
func CreateMultipleData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var students []models.StudentForm
	_ = json.NewDecoder(r.Body).Decode(&students)
	var data []interface{}
	for _, student := range students {
		data = append(data, student)
	}
	insertMultipleData(data)
	json.NewEncoder(w).Encode(students)
}

// update student data
func UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	studentId := r.URL.Query().Get("id")
	updateSingleData(studentId)
	json.NewEncoder(w).Encode("Data updated successfully")
}

// delete student data by id
func DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	studentId := r.URL.Query().Get("id")
	deleteSingleData(studentId)
	json.NewEncoder(w).Encode("Data deleted successfully")
}

// delete all student data
func DeleteAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAllData()
	json.NewEncoder(w).Encode("All data deleted successfully")
}
