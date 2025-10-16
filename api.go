package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CatFact struct {
	CatFact string `json:"fact"`
}

func makeRequest() (*CatFact, int, error) {
	client := &http.Client{
		Timeout: 5 * time.Second, // set 10 second timeout
	}

	resp, err := client.Get("https://catfact.ninja/fact")
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var fact CatFact
	if err := json.NewDecoder(resp.Body).Decode(&fact); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to decode response: %v", err)
	}

	return &fact, resp.StatusCode, nil
}
