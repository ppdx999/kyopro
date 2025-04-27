package atcoder_test

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ppdx999/kyopro/internal/infrastructure/external_service/atcoder"
)

func TestLoginCheck(t *testing.T) {
	tests := []struct {
		name          string
		reqErr        error
		resStatusCode int
		want          bool
		wantErr       bool
	}{
		{
			name:          "correct case",
			reqErr:        nil,
			resStatusCode: http.StatusOK,
			want:          true,
		},
		{
			name:          "request error",
			reqErr:        errors.New("request error"),
			resStatusCode: http.StatusOK,
			want:          false,
			wantErr:       true,
		},
		{
			name:          "invalid status code",
			reqErr:        nil,
			resStatusCode: http.StatusInternalServerError,
			want:          false,
			wantErr:       true,
		},
		{
			name:          "redirect to login page",
			reqErr:        nil,
			resStatusCode: http.StatusFound,
			want:          false,
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRequester := NewMockRequester(mockCtrl)
			statusCode := tt.resStatusCode
			if statusCode == 0 {
				statusCode = http.StatusOK
			}
			path := "/contests/abc001/submit"
			if statusCode == http.StatusFound {
				path = "/login"
			}
			mockRequester.EXPECT().Request(gomock.Any()).Return(&http.Response{
				Request: &http.Request{
					URL: &url.URL{
						Path: path,
					},
				},
				StatusCode: statusCode,
				Body:       io.NopCloser(strings.NewReader("")),
			}, tt.reqErr)

			atcoder := atcoder.NewAtcoder(mockRequester)

			got, err := atcoder.LoginCheck()

			if (err != nil) != tt.wantErr {
				t.Errorf("LoginCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("LoginCheck() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
