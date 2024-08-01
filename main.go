//go:build !codeanalysis
// +build !codeanalysis

package main

import (
	"log"
	"os"

	"github.com/mnestor/opnsense-go/opnsense"
)

func main() {

	// chang logging to a separate file
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	fwLogger := log.New(file, "", 0)
	opnsense.SetLogger(fwLogger)

	// now create the client
	_, err = opnsense.NewClient("http://localhost:8080", "", "", true)
	if err != nil {
		log.Fatal(err)
	}
}
