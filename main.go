package main

import (
	"fmt"

	router "github.com/rcgc/go-hexagonal.git/src/app"
)

func main() {
	r := router.NewRouter()
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("error starting server: ", err)
	}
}