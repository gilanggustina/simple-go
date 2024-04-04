package main

import {
	"fmt"
	"net/http"
	"os"
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	message := os.Getenv("HELLO_MESSAGE")
	if message == "" {
		message = "Hello, World! (not set env)"
	}
	fmt.Fprintf(w, message)
}

func main()  {
	http.HandleFunc("/", helloWorld)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running at http://localhost: %s\n" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n",err)
	}
}