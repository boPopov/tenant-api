package integration

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Base URL for Docker API
const baseURL = "http://localhost:3000"

// Mock Response for OAuth Provider (e.g., GitHub)
type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// Simulated OAuth 2.0 Login Flow
func TestOAuthFlowIntegration(t *testing.T) {
	// Step 1: Simulate OAuth Provider Response
	// oauthMockResponse := OAuthResponse{
	// 	AccessToken: mocks.OAuthMockGenerateToken("oauthIntegrationUser"),
	// 	TokenType:   "Bearer",
	// }

	// Step 2: Mock OAuth Callback (`/auth/callback`)
	// reqBody, _ := json.Marshal(oauthMockResponse) bytes.NewBuffer(reqBody)
	req, err := http.NewRequest(http.MethodGet, baseURL+"/api/auth/github/callback", nil)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Step 3: Decode Access Token
	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)
	log.Println(response["access_token"])

	accessToken, exists := response["access_token"].(string)
	assert.True(t, exists, "Expected an access token in the response")

	// Step 4: Use Token to Access Protected Route
	req2, err := http.NewRequest(http.MethodGet, baseURL+"/api/tenants", nil)
	assert.NoError(t, err)

	req2.Header.Set("Authorization", "Bearer "+accessToken)

	resp2, err := client.Do(req2)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp2.StatusCode)

	// Step 5: Decode and Assert Tenant List
	var tenants []map[string]interface{}
	json.NewDecoder(resp2.Body).Decode(&tenants)
	assert.Greater(t, len(tenants), 0, "Expected some tenants to be listed")
}
