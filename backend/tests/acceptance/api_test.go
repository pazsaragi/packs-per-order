package acceptance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// These tests are ran against the deployed environment
const baseURL = "https://packs-per-order.onrender.com"

type PackResponse struct {
	Packs map[string]int `json:"packs"`
}

func TestHealthCheck(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("%s/health-check", baseURL))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result map[string]string
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, "pong!", result["ping"])
}

func TestPackEndpoint(t *testing.T) {
	testCases := []struct {
		order    int
		expected PackResponse
	}{
		{250, PackResponse{Packs: map[string]int{"250": 1}}},
		{251, PackResponse{Packs: map[string]int{"500": 1}}},
		{501, PackResponse{Packs: map[string]int{"500": 1, "250": 1}}},
		{12001, PackResponse{Packs: map[string]int{"5000": 2, "2000": 1, "250": 1}}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Order: %d", tc.order), func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf("%s/pack?order=%d", baseURL, tc.order))
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var result PackResponse
			err = json.NewDecoder(resp.Body).Decode(&result)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected.Packs, result.Packs)
		})
	}
}
