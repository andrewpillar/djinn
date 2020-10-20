// +build integration

package integration

import (
	"net/http"
	"testing"
)

func Test_Variable(t *testing.T) {
	client := newClient(server)

	variable := map[string]interface{}{
		"key":   "PGADDR",
		"value": "host=localhost port=5432 dbname=djinn user=djinn password=secret sslmode=disable",
	}

	reqs := []request{
		{
			name:        "attempt to create variable",
			method:      "POST",
			uri:         "/api/variables",
			token:       myTok,
			contentType: "application/json",
			code:        http.StatusBadRequest,
		},
		{
			name:        "create variable",
			method:      "POST",
			uri:         "/api/variables",
			token:       myTok,
			contentType: "application/json",
			body:        jsonBody(variable),
			code:        http.StatusCreated,
		},
		{
			name:        "create variable with duplicate name",
			method:      "POST",
			uri:         "/api/variables",
			token:       myTok,
			contentType: "application/json",
			body:        jsonBody(variable),
			code:        http.StatusBadRequest,
		},
	}

	for _, req := range reqs {
		client.do(t, req)
	}

	client.do(t, request{
		name:        "search for variable with 'PG'",
		method:      "GET",
		uri:         "/api/variables?search=PG",
		token:       myTok,
		contentType: "application/json",
		code:        http.StatusOK,
		check:       checkResponseJSONLen(1),
	})
}
