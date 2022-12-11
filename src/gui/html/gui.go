package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

// CountState represents the current count state
type CountState struct {
	Count int
}

func main() {
	// Initialize the count state
	countState := CountState{0}

	// Create the HTML template for the GUI
	tmpl := template.Must(template.New("").Parse(`
		<html>
			<head>
				<title>Count Up</title>
			</head>
			<body>
				<h1>Count: {{.Count}}</h1>
				<button onclick="window.location.reload()">Count Up</button>
			</body>
		</html>
	`))

	// Handle requests to the root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Increment the count state by a random number in the range [1-10]
		countState.Count += rand.Intn(10) + 1

		// Execute the HTML template and write the result to the response writer
		if err := tmpl.Execute(w, countState); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
