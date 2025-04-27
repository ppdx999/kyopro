package atcoder_test

import (
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ppdx999/kyopro/internal/domain/model"
	"github.com/ppdx999/kyopro/internal/infrastructure/external_service/atcoder"
)

func TestGetProblemIds(t *testing.T) {
	tests := []struct {
		name          string
		html          string
		reqErr        error
		resStatusCode int
		want          []model.ProblemId
		wantErr       bool
	}{
		{
			name: "correct case",
			html: `<a href="/contests/abc100/tasks/abc100_a">Problem A</a>
			<a href="/contests/abc100/tasks/abc100_b">Problem B</a>
			<a href="/contests/abc100/tasks/abc100_c">Problem C</a>`,
			reqErr: nil,
			want:   []model.ProblemId{"a", "b", "c"},
		},
		{
			name: "no problem found",
			html: `<html><body></body></html>`,
			want: nil,
		},
		{
			name:    "request error",
			html:    "",
			reqErr:  errors.New("request error"),
			want:    nil,
			wantErr: true,
		},
		{
			name:          "invalid status code",
			html:          "",
			reqErr:        nil,
			resStatusCode: http.StatusInternalServerError,
			want:          nil,
			wantErr:       true,
		},
		{
			name:          "empty body",
			html:          "",
			reqErr:        nil,
			resStatusCode: http.StatusOK,
			want:          nil,
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requestUrl string
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRequester := NewMockRequester(mockCtrl)
			mockRequester.EXPECT().Request(gomock.Any()).
				DoAndReturn(func(req *http.Request) (*http.Response, error) {
					requestUrl = req.URL.String()
					statusCode := tt.resStatusCode
					if statusCode == 0 {
						statusCode = http.StatusOK
					}

					return &http.Response{
						StatusCode: statusCode,
						Body:       io.NopCloser(strings.NewReader(tt.html)),
					}, tt.reqErr
				},
				)

			atcoder := atcoder.NewAtcoder(mockRequester)

			got, err := atcoder.GetProblemIds("abc100")

			if (err != nil) != tt.wantErr {
				t.Errorf("GetProblemIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProblemIds() = %v, want %v", got, tt.want)
				return
			}

			if !reflect.DeepEqual(requestUrl, "https://atcoder.jp/contests/abc100/tasks") {
				t.Errorf("GetProblemIds() = %v, want %v", requestUrl, "https://atcoder.jp/contests/abc100/tasks")
				return
			}
		})
	}

}
