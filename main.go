package main

import (
	"fmt"
	"log"
	"net/http"

	err "github.com/viggy28/go-til/error"
	"github.com/viggy28/go-til/interfaces"
)

func main() {
	d := make(interfaces.DB)
	d["ssd"] = 100
	d["nvme"] = 1000
	http.ListenAndServe(":8080", d)
	child2 := err.Child2Error{}
	parent := fmt.Errorf("parent: %w", child2)
	grandparent := fmt.Errorf("grandparent: %w", parent)
	if err := err.CheckErrorLevel(grandparent); err != nil {
		log.Printf("Kid, please don't repeat it %v", err)
	}
	child1 := err.Child1Error{}
	parent = fmt.Errorf("parent: %w", child1)
	grandparent = fmt.Errorf("grandparent: %w", parent)
	err.CheckErrorLevel(grandparent)
	if err := err.CheckErrorLevel(grandparent); err != nil {
		log.Printf("Kid, please don't repeat it %v", err)
	}
}
