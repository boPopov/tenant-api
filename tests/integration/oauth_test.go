package integration

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

// Mock OAuth2 Token Server
func mockOAuthServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/token" {
			w.Header().Set("Content-Type", "application/json")
			// Return mock access token
			fmt.Fprintln(w, `{"access_token":"mockAccessToken","token_type":"Bearer"}`)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
}

// Test OAuth Integration
func TestOAuthFlowIntegration(t *testing.T) {
	// Step 1: Start Mock OAuth Server
	oauthServer := mockOAuthServer()
	defer oauthServer.Close()

	// Step 2: Use Mock OAuth URL in oauthConfig
	mockOAuthConfig := &oauth2.Config{
		ClientID:     "mock-client-id",
		ClientSecret: "mock-client-secret",
		Endpoint: oauth2.Endpoint{
			TokenURL: oauthServer.URL + "/token",
		},
	}

	// Simulate GitHub Callback (OAuth Code Exchange)
	token, err := mockOAuthConfig.Exchange(context.Background(), "mockCode")
	assert.NoError(t, err, "Failed to exchange mock code")
	assert.Equal(t, "mockAccessToken", token.AccessToken, "Expected mock access token")

	// Simulate API Call to Protected Route `/api/tenants`
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/api/tenants", nil)
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err, "Request to /api/tenants failed")

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected 200 OK from /api/tenants")

	body, _ := io.ReadAll(resp.Body)
	var tenants []map[string]interface{}
	json.Unmarshal(body, &tenants)
	assert.Greater(t, len(tenants), 0, "Expected tenant list")
}
