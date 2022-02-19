package main

import (
	no1datamanipulation "citcall-go-test/no1-data-manipulation"
	no3problemsolving "citcall-go-test/no3-problem-solving"
	"log"
)

func main() {
	// Problem Number 1 Make Table
	if err := no1datamanipulation.MakeTable(); err != nil {
		log.Println(err)
	}

	// Problem Number 3 Bebek
	no3problemsolving.Test()
}
