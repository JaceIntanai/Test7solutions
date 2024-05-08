package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func countMeat(text string) map[string]int {
	counts := make(map[string]int)
	beefs := strings.Fields(text)

	for _, beef := range beefs {
		beef = strings.ToLower(strings.Trim(beef, ",."))
		counts[beef]++
	}
	return counts
}

func beefSummary(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error parsing form data: %v", err)
		return
	}

	text := r.FormValue("text")
	// text := "Fatback t-bone t-bone, pastrami .. t-bone. pork, meatloaf jowl enim. Bresaola t-bone."

	counts := countMeat(text)
	delete(counts, "")

	// Create JSON response
	result := map[string]map[string]int{"beef": counts}
	jsonResponse, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error encoding JSON response:", err)
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send JSON response
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/beef/summary", beefSummary)
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
