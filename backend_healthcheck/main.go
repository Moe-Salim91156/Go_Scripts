package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func checkDjango() {

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client := &http.Client{}
	var Url string = "http://localhost:8000/health/"
	req, err := http.NewRequestWithContext(ctx, "GET", Url, nil)
	req.Header.Set("Accept", "application/json")
	if err != nil {
		log.Fatal("Error creating a request", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error Running Requeset", err)
	}
	defer resp.Body.Close()
	fmt.Printf("Status: %s\n", resp.Status)

	if resp.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Health Check Response: %s\n", string(body))
	}
}

// run a get request on django endpoint
// and return the result
