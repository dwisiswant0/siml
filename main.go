package main

import (
	"fmt"
	"os"

	"github.com/dwisiswant0/siml/pkg/siml"
)

func init() {
	if len(os.Args) < 2 {
		fmt.Printf(USAGE, VERSION)
		os.Exit(2)
	} else {
		domain = os.Args[1]
	}
}

func main() {
	similar, _ := siml.Get(domain)
	for _, s := range similar {
		fmt.Println(s)
	}
}
