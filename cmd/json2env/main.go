package main

import (
	"fmt"
	"json2env/internal/services/env"
	"json2env/internal/services/json"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: json2env <file.json> <file.env>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := json.ReadJson(inputFile)
	if err != nil {
		panic(err)
	}

	err = env.CreateEnvFile(outputFile, data)
	if err != nil {
		panic(err)
	}
}
