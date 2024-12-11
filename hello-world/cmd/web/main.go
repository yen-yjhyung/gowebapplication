package main

import (
	"fmt"
	"net/http"

	"github.com/yen-yjhyung/gowebapplication/hello-world/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
