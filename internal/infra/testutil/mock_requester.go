package testutil

import (
	"io"
	"net/http"
	"strings"
)

type MockRequester struct {
	ResponseStatusCode int
	ResponseBody       string
	ResponseErr        error
	CalledWith         *http.Request
}

func (m *MockRequester) Request(req *http.Request) (*http.Response, error) {
	m.CalledWith = req

	if m.ResponseErr != nil {
		return nil, m.ResponseErr
	}

	statusCode := http.StatusOK
	if m.ResponseStatusCode != 0 {
		statusCode = m.ResponseStatusCode
	}

	resp := &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(strings.NewReader(m.ResponseBody)),
		Request:    req,
		Header:     make(http.Header),
	}
	return resp, nil
}
