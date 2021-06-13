package interfaces

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	d := make(DB)
	d["ssd"] = 100
	d["nvme"] = 1000
	// ListenAndServer needs a handler (to handle requests)
	http.ListenAndServe(":8080", d)
}

// type DB satisfies http.Handler
type DB map[string]uint

func (d DB) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	for k, v := range d {
		// Fprintf takes io.writer as the first argument. Since http.ResponseWriter also satisfies io.writer we can
		// use that
		n, err := fmt.Fprintf(w, "%s: %d\n", k, v)
		if err != nil {
			log.Printf("err %v", err)
		}
		log.Println("INF: ", n)
	}
	return
}
