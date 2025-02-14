package unittests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"net/http"

	mocks "github.com/boPopov/tenant-api/tests/mock"
	"github.com/stretchr/testify/assert"
)

var bearer_token string
var baseUrl string
var newTenantId int

func init() {
	bearer_token = mocks.MockGenerateJWT("test")
	baseUrl = "http://localhost:3000/api"
}

/*
*
Function makeApiCall is executing the API calls for the tests. It requires the method, apiRoute to be send.
If there are no queryParameters an empty string is accepted, also if there is no Body for the request an empty string is accepted.
*/
func makeApiCall(t *testing.T, method string, apiRoute string, queryParameter string, requestBody string, unauthorized bool) (*http.Response, error) {
	var body io.Reader
	if requestBody != "" {
		body = bytes.NewBuffer([]byte(requestBody))
	}
	// log.Println("In makeAPICAll")
	// log.Println("Auth token is: ")
	// log.Println(bearer_token)

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s%s", baseUrl, apiRoute, queryParameter), body)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	if !unauthorized {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearer_token))
	} else {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %sunauthorized", bearer_token))
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

/*
*
parseTenantId function, is used for setting the variable @newTenantId.
After the TestCreateTenantSuccess Unit test is executed the API returns all of the entered data, plus the ID.
This ID is later needed for all of the Test*Success Unit tests.
*/
func parseTenantId(response *http.Response) {
	var responseMap map[string]interface{}

	if err := json.NewDecoder(response.Body).Decode(&responseMap); err != nil {
		return
	}

	tenantID, ok := responseMap["ID"].(float64)
	if !ok {
		return
	}

	newTenantId = int(tenantID)
}

func TestCreateTenantSuccess(t *testing.T) {
	body := `{"name": "testtt", "email": "testt@tes.com", "active": true }`
	resp, err := makeApiCall(t, http.MethodPost, "/tenants", "", body, false)

	parseTenantId(resp)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestGetAllTenantSuccess(t *testing.T) {
	resp, err := makeApiCall(t, http.MethodGet, "/tenants", "", "", false)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetTenantSuccess(t *testing.T) {
	resp, err := makeApiCall(t, http.MethodGet, "/tenants", fmt.Sprintf("/%d", newTenantId), "", false)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdateTenantSuccess(t *testing.T) {
	body := `{"name": "Test", "email": "test@gmail.com", "active": true }`
	resp, err := makeApiCall(t, http.MethodPut, "/tenants", fmt.Sprintf("/%d", newTenantId), body, false)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeleteTenantSuccess(t *testing.T) {
	resp, err := makeApiCall(t, http.MethodDelete, "/tenants", fmt.Sprintf("/%d", newTenantId), "", false)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestCreateTenantsUnauthorized(t *testing.T) {
	body := `{"name": "testtt", "email": "testt@tes.com", "active": true }`
	resp, err := makeApiCall(t, http.MethodPost, "/tenants", "", body, true)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestGetAllTenantsUnauthorized(t *testing.T) {
	resp, err := makeApiCall(t, http.MethodGet, "/tenants", "", "", true)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestGetTenantUnauthorized(t *testing.T) {
	resp, err := makeApiCall(t, http.MethodGet, "/tenants", "/1", "", true)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestUpdateTenantUnauthorized(t *testing.T) {
	body := `{"name": "Test", "email": "test@gmail.com", "active": true }`
	resp, err := makeApiCall(t, http.MethodPut, "/tenants", "/1", body, true)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestDeleteTenantUnauthorized(t *testing.T) {
	resp, err := makeApiCall(t, http.MethodDelete, "/tenants", "/1", "", true)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestUpdateTenantNotFound(t *testing.T) {
	body := `{"name": "Test", "email": "test@gmail.com", "active": true }`
	resp, err := makeApiCall(t, http.MethodPut, "/tenants", "/150", body, false)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGetTenantNotFound(t *testing.T) {
	resp, err := makeApiCall(t, http.MethodGet, "/tenants", "/150", "", false)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDeleteTenantNotFound(t *testing.T) {
	resp, err := makeApiCall(t, http.MethodDelete, "/tenants", "/150", "", false)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestCreateTenantIvalidData(t *testing.T) {
	resp, err := makeApiCall(t, http.MethodPost, "/tenants", "", "", false)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
