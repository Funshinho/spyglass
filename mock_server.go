package spyglass

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// GetMockServer initializes the mock responses when calling data api
func GetMockServer(t *testing.T) *httptest.Server {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if strings.Contains(req.URL.Path, "teams/TOR") {
			writeContent(t, &w, "2023_players_TOR.xml")
		}
		if strings.Contains(req.URL.Path, "teams") {
			writeContent(t, &w, "2023_teams.xml")
		}
	}))
	return server
}

func writeContent(t *testing.T, w *http.ResponseWriter, filename string) {
	content, err := ioutil.ReadFile("mocks/" + filename)
	if err != nil {
		t.Fatalf("Could not read file: " + filename)
	}
	(*w).WriteHeader(http.StatusOK)
	(*w).Write(content)
}
