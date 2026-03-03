package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Run("reads token from env", func(t *testing.T) {
		token := "test-token-123"
		os.Setenv("GITHUB_TOKEN", token)
		defer os.Unsetenv("GITHUB_TOKEN")

		client := NewClient()
		if client.token != token {
			t.Errorf("NewClient() token = %q, want %q", client.token, token)
		}
	})
}

func TestRequest(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		setupServer    func(*httptest.Server)
		wantStatusCode int
		wantError      bool
	}{
		{
			name:   "GET request success",
			method: http.MethodGet,
			setupServer: func(s *httptest.Server) {
				s.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.Method != http.MethodGet {
						t.Errorf("expected GET, got %s", r.Method)
					}
					if r.Header.Get("Authorization") != "Bearer test-token" {
						t.Error("missing or wrong Authorization header")
					}
					if r.Header.Get("Accept") != "application/vnd.github+json" {
						t.Error("missing or wrong Accept header")
					}
					if r.Header.Get("X-GitHub-Api-Version") != "2022-11-28" {
						t.Error("missing or wrong X-GitHub-Api-Version header")
					}
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"data":"test"}`))
				})
			},
			wantStatusCode: http.StatusOK,
			wantError:      false,
		},
		{
			name:   "POST request includes Content-Type",
			method: http.MethodPost,
			setupServer: func(s *httptest.Server) {
				s.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.Header.Get("Content-Type") != "application/json" {
						t.Error("missing or wrong Content-Type header")
					}
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte(`{"created":true}`))
				})
			},
			wantStatusCode: http.StatusCreated,
			wantError:      false,
		},
		{
			name:   "404 error returns error",
			method: http.MethodGet,
			setupServer: func(s *httptest.Server) {
				s.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNotFound)
					w.Write([]byte(`{"error":"not found"}`))
				})
			},
			wantStatusCode: http.StatusNotFound,
			wantError:      true,
		},
		{
			name:   "500 error returns error",
			method: http.MethodGet,
			setupServer: func(s *httptest.Server) {
				s.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
				})
			},
			wantStatusCode: http.StatusInternalServerError,
			wantError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(nil)
			tt.setupServer(server)

			client := NewClientWithBaseURL(server.URL)
			client.token = "test-token"
			body, headers, err := client.request(tt.method, server.URL, nil)

			if tt.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !tt.wantError {
				if headers == nil {
					t.Error("expected headers, got nil")
				}
				if body == nil {
					t.Error("expected body, got nil")
				}
			}

			server.Close()
		})
	}
}

func TestGetAssignedOpenIssues(t *testing.T) {
	tests := []struct {
		name        string
		setupServer func(*httptest.Server)
		wantLength  int
		wantError   bool
	}{
		{
			name: "single page of issues",
			setupServer: func(s *httptest.Server) {
				s.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Query().Get("per_page") != "100" {
						t.Error("missing or wrong per_page parameter")
					}
					issues := []GetIssuesResponse{
						{Number: 1},
						{Number: 2},
					}
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(issues)
				})
			},
			wantLength: 2,
			wantError:  false,
		},
		{
			name: "multiple pages of issues",
			setupServer: func(s *httptest.Server) {
				callCount := 0
				s.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					callCount++
					if callCount == 1 {
						issues := []GetIssuesResponse{
							{Number: 1},
						}
						w.Header().Set("Content-Type", "application/json")
						w.Header().Set("Link", fmt.Sprintf(`<%s>; rel="next"`, s.URL+"?page=2"))
						w.WriteHeader(http.StatusOK)
						json.NewEncoder(w).Encode(issues)
					} else if callCount == 2 {
						issues := []GetIssuesResponse{
							{Number: 2},
							{Number: 3},
						}
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(http.StatusOK)
						json.NewEncoder(w).Encode(issues)
					}
				})
			},
			wantLength: 3,
			wantError:  false,
		},
		{
			name: "empty response",
			setupServer: func(s *httptest.Server) {
				s.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("[]"))
				})
			},
			wantLength: 0,
			wantError:  false,
		},
		{
			name: "network error",
			setupServer: func(s *httptest.Server) {
				s.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(`{"message":"internal error"}`))
				})
			},
			wantLength: 0,
			wantError:  true,
		},
		{
			name: "invalid json response",
			setupServer: func(s *httptest.Server) {
				s.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`invalid json`))
				})
			},
			wantLength: 0,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(nil)
			tt.setupServer(server)

			client := NewClientWithBaseURL(server.URL)
			client.token = "test-token"
			issues, err := client.GetAssignedOpenIssues()

			if tt.wantError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !tt.wantError && len(issues) != tt.wantLength {
				t.Errorf("got %d issues, want %d", len(issues), tt.wantLength)
			}

			server.Close()
		})
	}
}
