package hisend

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Request(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-api-key" {
			t.Errorf("Expected token 'Bearer test-api-key', got '%s'", r.Header.Get("Authorization"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type 'application/json', got '%s'", r.Header.Get("Content-Type"))
		}
		if r.URL.Path != "/domains" {
			t.Errorf("Expected path '/domains', got '%s'", r.URL.Path)
		}
		
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode([]Domain{{ID: 1, Name: "example.com"}})
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	client := NewClient(Config{APIKey: "test-api-key"})
	client.BaseURL = server.URL // Override BaseURL for testing

	domains, err := client.Domains.List()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(domains) != 1 {
		t.Fatalf("Expected 1 domain, got %d", len(domains))
	}
	if domains[0].Name != "example.com" {
		t.Errorf("Expected domain name 'example.com', got '%s'", domains[0].Name)
	}
}
