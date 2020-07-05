package codinggoweb

import (
	"fmt"
	"net/http"
)

// RunBasicHTTP : function to run a basic http server
func RunBasicHTTP() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you have requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
