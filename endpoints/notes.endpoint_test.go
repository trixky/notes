package endpoints

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNotes(t *testing.T) {
	const target = "http://localhost:3000/notes"

	tests := []struct {
		Method         string
		ExpectedStatus int
		Body           string
	}{
		// ---------------------------- GET
		{
			Method:         "GET",
			ExpectedStatus: 200,
			Body:           "",
		},
		// ---------------------------- POST
		{
			Method:         "POST",
			ExpectedStatus: 400,
			Body:           "",
		},
		{
			Method:         "POST",
			ExpectedStatus: 400,
			Body:           "{\"message\": \"\"}",
		},
		{
			Method:         "POST",
			ExpectedStatus: 400,
			Body:           "{\"message\": \"\", \"tag\": \"TaG\"}",
		},
		{
			Method:         "POST",
			ExpectedStatus: 200,
			Body:           "{\"message\": \"123\"}",
		},
		{
			Method:         "POST",
			ExpectedStatus: 200,
			Body:           "{\"message\": \"123\", \"tag\": \"TaG\"}",
		},
		// ---------------------------- DELETE
		{
			Method:         "DELETE",
			ExpectedStatus: 200,
			Body:           "",
		},
		// ---------------------------- PUT [not allowed]
		{
			Method:         "PUT",
			ExpectedStatus: http.StatusMethodNotAllowed,
			Body:           "",
		},
		// ---------------------------- PATCH [not allowed]
		{
			Method:         "PATCH",
			ExpectedStatus: http.StatusMethodNotAllowed,
			Body:           "",
		},
		// ---------------------------- COPY [not allowed]
		{
			Method:         "COPY",
			ExpectedStatus: http.StatusMethodNotAllowed,
			Body:           "",
		},
		// ---------------------------- HEAD [not allowed]
		{
			Method:         "HEAD",
			ExpectedStatus: http.StatusMethodNotAllowed,
			Body:           "",
		},
		// ---------------------------- OPTIONS [not allowed]
		{
			Method:         "OPTIONS",
			ExpectedStatus: http.StatusMethodNotAllowed,
			Body:           "",
		},
		// ---------------------------- LINK [not allowed]
		{
			Method:         "LINK",
			ExpectedStatus: http.StatusMethodNotAllowed,
			Body:           "",
		},
		// ---------------------------- UNLINK [not allowed]
		{
			Method:         "UNLINK",
			ExpectedStatus: http.StatusMethodNotAllowed,
			Body:           "",
		},
		// ---------------------------- PURGE [not allowed]
		{
			Method:         "PURGE",
			ExpectedStatus: http.StatusMethodNotAllowed,
			Body:           "",
		},
		// ---------------------------- LOCK [not allowed]
		{
			Method:         "LOCK",
			ExpectedStatus: http.StatusMethodNotAllowed,
			Body:           "",
		},
		// ---------------------------- UNLOCK [not allowed]
		{
			Method:         "UNLOCK",
			ExpectedStatus: http.StatusMethodNotAllowed,
			Body:           "",
		},
		// ---------------------------- PROPFIND [not allowed]
		{
			Method:         "PROPFIND",
			ExpectedStatus: http.StatusMethodNotAllowed,
			Body:           "",
		},
		// ---------------------------- VIEW [not allowed]
		{
			Method:         "VIEW",
			ExpectedStatus: http.StatusMethodNotAllowed,
			Body:           "",
		},
	}

	for _, test := range tests {
		bodyReader := strings.NewReader(test.Body)

		req := httptest.NewRequest(test.Method, target, bodyReader)
		w := httptest.NewRecorder()
		Notes(w, req)

		resp := w.Result()

		if resp.StatusCode != test.ExpectedStatus {
			t.Fatalf("status code don't match: %v | result: %v", test.ExpectedStatus, resp.StatusCode)
		}
	}
}
