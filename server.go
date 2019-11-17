package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./static/")))

	http.HandleFunc("/increment", incrementCounter)

	log.Fatal(http.ListenAndServe(":8081", nil))

}
