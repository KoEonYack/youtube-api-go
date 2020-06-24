package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elliotforbes/youtube-stats/youtube"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Youtbue Subscriber Monitor")

	item, err := youtube.GetSubscribers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", item)

	//setupRoutes()
}
