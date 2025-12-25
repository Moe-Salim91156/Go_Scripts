package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type HealthResult struct {
	StatusCode int
	Body       string
	Err        error
}

func checkHealth(Url string, client *http.Client) HealthResult {

	ctx, cancel := context.WithTimeout(context.Background(), client.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, Url, nil)
	if err != nil {
		return HealthResult{Err: fmt.Errorf("Shitty Request Havent Been Created Mate %w\n ", err)}
	}
	resp, err := client.Do(req)
	if err != nil {
		return HealthResult{Err: fmt.Errorf("Shitty Client Didn do the Request Mate %w\n ", err)}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return HealthResult{Err: fmt.Errorf("Reading Body Failed %w\n", err)}
	}
	return HealthResult{
		StatusCode: resp.StatusCode,
		Body:       string(body),
		Err:        nil,
	}
}

func createNewClient(timeout time.Duration) *http.Client {
	client := &http.Client{
		Timeout: timeout,
	}
	return (client)

}

func main() {
	client := createNewClient(2 * time.Second)
	Url := "http://localhost:8000/health/" // could expand it to something more modualr, but for now this works fine
	result := checkHealth(Url, client)
	if result.Err != nil {
		log.Fatalf("Health check error : %v", result.Err)
	}
	fmt.Println("Health Check Done Successfully")
	fmt.Printf("Status %d\n", result.StatusCode)
	fmt.Printf("Body: %s\n", result.Body)
}
