package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"json2env/internal/services/env"
	"json2env/internal/services/json"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: json2env <file.env>")
		return
	}
	outputFile := os.Args[1]

	jsonData, err := clipboard.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.ReadJson(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	err = env.CreateEnvFile(outputFile, data)
	if err != nil {
		log.Fatal(err)
	}
}
