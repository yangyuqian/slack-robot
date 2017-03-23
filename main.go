package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.Handle("/xiabeebee", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`You happy now?`))
	}))

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
