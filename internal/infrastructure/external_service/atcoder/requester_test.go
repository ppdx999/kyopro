package atcoder_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ppdx999/kyopro/internal/domain/model"
	"github.com/ppdx999/kyopro/internal/infrastructure/external_service/atcoder"
)

type MockSessionLoader struct {
	session string
}

func (m *MockSessionLoader) LoadSession() (model.SessionSecret, error) {
	return model.SessionSecret(m.session), nil
}

func (m *MockSessionLoader) SetSession(session string) {
	m.session = session
}

func TestRequest(t *testing.T) {
	tests := []struct {
		name       string
		resBody    string
		statusCode int
	}{
		{
			name:       "Correct case",
			resBody:    "OK",
			statusCode: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "%s", tt.resBody)
			})
			ts := httptest.NewServer(handler)
			defer ts.Close()

			loader := &MockSessionLoader{session: "session"}
			r := atcoder.NewAtcoderRequester(loader)

			req, _ := http.NewRequest("GET", ts.URL, nil)

			res, err := r.Request(req)
			if err != nil {
				t.Fatal(err)
			}

			if res.StatusCode != tt.statusCode {
				t.Errorf("expected status code %d, but got %d", tt.statusCode, res.StatusCode)
			}
		})
	}
}
