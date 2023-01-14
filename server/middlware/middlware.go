package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/acceleraterA/parking_booking_system/server/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// collection object/instance
var collection *mongo.Collection

// create connection with mongo db
func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func createDBInstance() {
	// DB connection string
	connectionString := os.Getenv("DB_URI")

	// Database Name
	dbName := os.Getenv("DB_NAME")

	// Collection name
	collName := os.Getenv("DB_COLLECTION_NAME")

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
}

// Createspot create spot route
func Createspot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "spot")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var spot models.ParkingSpot
	_ = json.NewDecoder(r.Body).Decode(&spot)
	// fmt.Println(spot, r.Body)
	insertOneSpot(spot)
	json.NewEncoder(w).Encode(spot)
}

// Deletespot delete one spot
func DeleteSpot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOnespot(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	// json.NewEncoder(w).Encode("spot not found")

}

func insertOneSpot(spot models.ParkingSpot) {
	insertResult, err := collection.InsertOne(context.Background(), spot)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

// delete one spot from the DB, delete by ID
func deleteOnespot(spot string) {
	fmt.Println(spot)
	id, _ := primitive.ObjectIDFromHex(spot)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}
