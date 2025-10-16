package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
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

	url := fmt.Sprintf("https://catfact.ninja/fact?max_length=%d", RandomInt(20, 170))

	resp, err := client.Get(url)
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

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
