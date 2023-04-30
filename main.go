package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dwisiswant0/siml/pkg/siml"
)

func init() {
	if len(os.Args) < 2 {
		fmt.Printf(usage, version)
		os.Exit(2)
	} else {
		domains = os.Args[1:]
	}
}

func main() {
	similar, err := siml.Gets(domains...)
	if err != nil {
		log.Fatal(err)
	}

	printed := make(map[string]bool)
	for _, s := range similar {
		if printed[s] {
			continue
		}

		fmt.Println(s)
		printed[s] = true
	}
}
