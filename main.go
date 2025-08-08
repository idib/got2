package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hjson/hjson-go/v4"
	"github.com/idib/got2/pkg/client"
)

func main() {
	// Create a new client
	c := client.NewClient("http://localhost:8080")

	fmt.Println("Sending ping request to server...")

	// Try to ping the server
	response, err := c.Ping()
	if err != nil {
		log.Fatalf("Failed to ping server: %v", err)
	}

	fmt.Printf("Server response: %s\n", response)

	// Example of how to use the client in a loop
	for i := 1; i <= 3; i++ {
		fmt.Printf("Ping attempt %d: ", i)

		response, err := c.Ping()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Success: %s\n", response)
		}

		time.Sleep(1 * time.Second)
	}
	sampleText := []byte(`
    {
        # specify rate in requests/second
        rate: 1000
        array:
        [
            foo
            bar
        ]
    }`)

	// We need to provide a variable where Hjson
	// can put the decoded data.
	var dat map[string]interface{}

	// Decode with default options and check for errors.
	if err := hjson.Unmarshal(sampleText, &dat); err != nil {
		panic(err)
	}
	// short for:
	// options := hjson.DefaultDecoderOptions()
	// err := hjson.UnmarshalWithOptions(sampleText, &dat, options)
	fmt.Println(dat)

	// In order to use the values in the decoded map,
	// we'll need to cast them to their appropriate type.

	rate := dat["rate"].(float64)
	fmt.Println(rate)

	array := dat["array"].([]interface{})
	str1 := array[0].(string)
	fmt.Println(str1)

	// To encode to Hjson with default options:
	sampleMap := map[string]int{"apple": 5, "lettuce": 7}
	hjson, _ := hjson.Marshal(sampleMap)
	// short for:
	// options := hjson.DefaultOptions()
	// hjson, _ := hjson.MarshalWithOptions(sampleMap, options)
	fmt.Println(string(hjson))
}
