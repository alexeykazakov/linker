package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("Open http://localhost:8090 in browser...")
	http.ListenAndServe(":8090", nil)
}
