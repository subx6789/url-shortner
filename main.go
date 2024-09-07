package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

type URL struct {
	Id           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

// In-memory database to store original and short URLs
var urlDB = make(map[string]URL)

func main() {
	// Register the root handler and specific handlers for shortening and redirecting URLs
	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", shortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)
	// Start the server on port 8080 and listen for incoming HTTP requests
	fmt.Println("Server starting on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// Log any error if the server fails to start
		fmt.Println("Error starting server :", err)
	}
}

// Root handler function for the base URL "/"
func handler(w http.ResponseWriter, r *http.Request) {
	// Send a welcome message as the response
	fmt.Fprintf(w, "Welcome to the URL Shortner")
}

// Handler for shortening a URL
func shortURLHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Parse and decode the JSON request body to extract the original URL
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		// If the request body is invalid, return a bad request error
		http.Error(w, "Invalid request Body", http.StatusBadRequest)
		return
	}
	// Generate the short URL from the original URL
	shortURL := createURL(data.URL)
	// Send the generated short URL as a JSON response
	response := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: shortURL,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler for redirecting to the original URL using the short URL
func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Extract the short URL from the URL path
	shortURL := strings.TrimPrefix(r.URL.Path, "/redirect/")
	// Retrieve the original URL using the short URL from the in-memory database
	url, err := getURL(shortURL)
	if err != nil {
		// If the short URL is not found, return a not found error
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	// Redirect to the original URL
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

// Function to retrieve the original URL using the short URL
func getURL(shortURL string) (URL, error) {
	// Look up the short URL in the in-memory database
	url, ok := urlDB[shortURL]
	// If the short URL is not found, return an error
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	// Return the found URL and no error
	return url, nil
}

// Function to create a short URL and store it in the in-memory database
func createURL(OriginalURL string) string {
	// Generate a short URL from the original URL using the MD5 hashing algorithm
	shortURL := generateShortURL(OriginalURL)
	// Generate a unique identifier using the "uuidgen" command
	id_byte, _ := exec.Command("uuidgen").Output()
	// Convert the byte slice to a string representation of the unique ID
	id := strings.TrimSpace(string(id_byte)) // Trim newline characters
	// Store the original URL, short URL, and other metadata in the database
	urlDB[shortURL] = URL{
		Id:           id,          // Unique identifier (UUID)
		OriginalURL:  OriginalURL, // The original URL provided by the user
		ShortURL:     shortURL,    // The generated short URL
		CreationDate: time.Now(),  // The timestamp when the short URL was created
	}
	// Return the generated short URL
	return shortURL
}

// Function to generate a short URL from the original URL using MD5 hashing
func generateShortURL(OriginalURL string) string {
	// Create a new MD5 hash object
	hasher := md5.New()
	// Write the original URL as a byte slice into the MD5 hasher
	hasher.Write([]byte(OriginalURL))
	// Compute the MD5 hash and store the result in a byte array
	data := hasher.Sum(nil)
	// Convert the byte array (MD5 hash) into a hexadecimal string
	hash := hex.EncodeToString(data)
	// Return the first 8 characters of the MD5 hash as the short URL
	return hash[:8]
}
