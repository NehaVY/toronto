package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Database credentials and configuration
const (
	DBUser     = "root"         // Replace with your MySQL username
	DBPassword = "Panda@sep18!" // Replace with your MySQL password
	DBName     = "toronto_time" // Database name you created
)

var db *sql.DB

// TimeResponse struct to send JSON response
type TimeResponse struct {
	CurrentTime string `json:"current_time"`
}

func main() {
	// Create a log file and set up logging
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()

	// Set log output to the file
	log.SetOutput(logFile)

	// Initialize database connection
	var dbErr error
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", DBUser, DBPassword, DBName)
	db, dbErr = sql.Open("mysql", dsn)
	if dbErr != nil {
		log.Fatalf("Failed to open MySQL database connection: %v", dbErr)
	}
	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	log.Println("MySQL database Connection done!")

	// Set up HTTP server
	http.HandleFunc("/current-time", handleCurrentTime)
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handleCurrentTime handles requests to /current-time endpoint
func handleCurrentTime(w http.ResponseWriter, r *http.Request) {
	// Log the incoming request
	log.Printf("Received request for current time from %s", r.RemoteAddr)

	// Get current time in Toronto (time zone handling)
	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		log.Printf("Error loading time zone: %v", err)
		http.Error(w, "Unable to load Toronto time zone", http.StatusInternalServerError)
		return
	}

	// Get the current time in Toronto's time zone
	currentTime := time.Now().In(loc)

	// Log the current time
	log.Printf("Current Toronto time: %v", currentTime)

	// Insert current time into the database
	_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", currentTime)
	if err != nil {
		log.Printf("Error inserting time into database: %v", err)
		http.Error(w, "Database insertion error", http.StatusInternalServerError)
		return
	}

	// Log successful insertion
	log.Printf("Current time is inserted: %v into the database", currentTime)

	// Create JSON response with the current time
	response := TimeResponse{
		CurrentTime: currentTime.Format(time.RFC1123),
	}

	// Set the Content-Type header and return the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error in JSON response: %v", err)
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
	}
}
