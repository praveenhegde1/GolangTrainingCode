package handlers
 
import (
    "LABORIS/Registration/Clients/db"
    "context"
    "encoding/json"
    "fmt"
 
    "LABORIS/Registration/Clients/models"
    "log"
    "net/http"
    "time"
 
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
 
var uri = "mongodb://localhost:27017"
 
func GetUrl() string {
 
    return uri
}
 
func GetClientHandler(w http.ResponseWriter, r *http.Request) {
 
    dbconn, err := db.Getconnection(uri)
 
    // Ping the MongoDB server to check the connection
 
    if err != nil {
        log.Fatal(err)
 
    }
 
    // Access the desired collection
    //collection := client.Database("Users").Collection("Profiles")
    //collection := dbconn.Database("golang").Collection("go")
    collection := dbconn.Database("Registration").Collection("Clients")
    // Fetch all documents from the collection
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
 
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        log.Fatal(err)
    }
 
    defer cursor.Close(ctx)
 
    // Iterate over the cursor and print the documents
    for cursor.Next(ctx) {
        var result bson.M
        err := cursor.Decode(&result)
        if err != nil {
            log.Fatal(err)
        }
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(result)
        fmt.Println(result)
    }
 
    if err := cursor.Err(); err != nil {
        log.Fatal(err)
    }
 
    // Disconnect from MongoDB
    err = dbconn.Disconnect(context.Background())
    if err != nil {
        log.Fatal(err)
    }
}
 
func UpdateClientHandler(w http.ResponseWriter, r *http.Request) {
    // Ensure the method is PUT
    if r.Method != http.MethodPut {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
 
    // Decode the JSON body into a User struct
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid input data", http.StatusBadRequest)
        return
    }
 
    // Validate that the ID is provided
    if user.ID == 0 {
        http.Error(w, "User ID is required", http.StatusBadRequest)
        return
    }
 
    dbconn, err := db.Getconnection(uri)
 
    // Access the "Profiles" collection
    collection := dbconn.Database("Registration").Collection("Clients")
 
    // Create the filter to find the user by ID
    filter := bson.M{"id": user.ID}
 
    // Define the update data
    update := bson.M{
        "$set": bson.M{
            "name":     user.Name,
            "email":    user.Email,
            "password": user.Password,
            "phone":    user.Phone,
            "address":  user.Address,
            "task":     user.TaskRequirements,
        },
    }
 
    // Update the user document in the collection
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
 
    result, _ := collection.UpdateOne(ctx, filter, update)
    if err != nil {
        http.Error(w, "Failed to update user", http.StatusInternalServerError)
        return
    }
 
    // Check if a document was modified
    if result.ModifiedCount == 0 {
        http.Error(w, "No user found with the given ID", http.StatusNotFound)
        return
    }
 
    // Respond with the updated user
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}
 
// AddUserHandler handles the HTTP request for adding a new user
func AddClientHandler(w http.ResponseWriter, r *http.Request) {
    // Only accept POST requests
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
 
    // Decode the JSON body into a User struct
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid input data", http.StatusBadRequest)
        return
    }
 
    dbconn, _ := db.Getconnection(uri)
    // Access the "Profiles" collection
    collection := dbconn.Database("Registration").Collection("Clients")
 
    // Insert the user into the collection
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
 
    _, err = collection.InsertOne(ctx, user)
    if err != nil {
        http.Error(w, "Failed to insert user into database", http.StatusInternalServerError)
        return
    }
 
    // Respond with the created user as JSON
    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
 
}
 
func DeleteClientHandler(w http.ResponseWriter, r *http.Request) {
    // Ensure the method is DELETE
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
 
    // Decode the JSON body into a User struct
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid input data", http.StatusBadRequest)
        return
    }
 
    // Validate that the ID is provided
    if user.ID == 0 {
        http.Error(w, "User ID is required", http.StatusBadRequest)
        return
    }
 
    dbconn, _ := db.Getconnection(uri)
    // Access the "Profiles" collection
    collection := dbconn.Database("Registration").Collection("Clients")
 
    // Define the filter to find the user by ID
    filter := bson.M{"id": user.ID}
 
    // Delete the user document from the collection
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
 
    result, err := collection.DeleteOne(ctx, filter)
    if err != nil {
        http.Error(w, "Failed to delete user", http.StatusInternalServerError)
        return
    }
 
    // Check if a document was deleted
    if result.DeletedCount == 0 {
        http.Error(w, "No user found with the given ID", http.StatusNotFound)
        return
    }
 
    // Respond with a success message
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "User deleted successfully",
    })
}
 
func LoginClientHandler(w http.ResponseWriter, r *http.Request) {
 
    // MongoDB connection string
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
 
    // Connect to MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
    defer func() {
        if err := client.Disconnect(context.TODO()); err != nil {
            log.Fatalf("Failed to disconnect MongoDB client: %v", err)
        }
    }()
 
    log.Println("Connected to MongoDB!")
    // Access the database and collection
    collection := client.Database("Registration").Collection("Clients")
 
    // Ensure the request method is POST
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
 
    // Parse the JSON body
    var userInput models.User
 
    err = json.NewDecoder(r.Body).Decode(&userInput)
    if err != nil {
        http.Error(w, "Invalid JSON body", http.StatusBadRequest)
        return
    }
 
    // Validate input
    if userInput.Phone == "" || userInput.Password == "" {
        http.Error(w, "Phone and password are required", http.StatusBadRequest)
        return
    }
 
    // // Define phone and password to check
    // phone := "1234567890"
    // password := "yourPassword"
 
    // Query to match the record
    filter := bson.M{"phone": userInput.Phone, "password": userInput.Password}
 
    var result models.User
 
    // Search in the collection
    err = collection.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            fmt.Println("Invalid phone or password.")
        } else {
            log.Fatalf("Error checking records: %v", err)
        }
        return
    }
 
    // If record matches
    fmt.Println("Login successful!")
 
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Login successful!"})
 
}
 
// HTTP handler function
func FindUserHandler(w http.ResponseWriter, r *http.Request) {
    // Ensure the request method is POST
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
 
    // Parse the JSON body
    var userInput models.User
 
    err := json.NewDecoder(r.Body).Decode(&userInput)
    if err != nil {
        http.Error(w, "Invalid JSON body", http.StatusBadRequest)
        return
    }
 
    // Validate input
    if userInput.Phone == "" && userInput.Email == "" {
        http.Error(w, "Phone or email  is required", http.StatusBadRequest)
        return
    }
    dbconn, _ := db.Getconnection(uri)
    // Set up MongoDB context
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
 
    collection := dbconn.Database("Registration").Collection("Clients")
 
    // Build the filter
    filter := bson.M{
        "$or": []bson.M{
            {"phone": userInput.Phone},
            {"email": userInput.Email},
        },
    }
 
    // Find matching documents
    cursor, err := collection.Find(ctx, filter)
    if err != nil {
        http.Error(w, "Error querying database: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)
 
    // Decode results
    var users []models.User
    for cursor.Next(ctx) {
        var user models.User
        if err := cursor.Decode(&user); err != nil {
            log.Println("Error decoding user:", err)
            continue
        }
        users = append(users, user)
    }
 
    if err := cursor.Err(); err != nil {
        http.Error(w, "Error processing results: "+err.Error(), http.StatusInternalServerError)
        return
    }
 
    // Return results as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}
 