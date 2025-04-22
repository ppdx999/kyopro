package atcoder_test

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/ppdx999/kyopro/internal/atcoder"
	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/testutil"
)

func TestExtractProblemIds(t *testing.T) {
	tests := []struct {
		name string
		html string
		want []model.ProblemId
	}{
		{
			name: "correct case",
			html: `<a href="/contests/abc100/tasks/abc100_a">Problem A</a>
		 		   <a href="/contests/abc100/tasks/abc100_b">Problem B</a>
		 		   <a href="/contests/abc100/tasks/abc100_c">Problem C</a>`,
			want: []model.ProblemId{"a", "b", "c"},
		},
		{
			name: "no problem found",
			html: `<html><body></body></html>`,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			atcoder := atcoder.NewAtcoder(&testutil.MockRequester{})

			got := atcoder.ExtractProblemIds(tt.html)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractProblemIds() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

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
			name:   "correct case",
			html:   `<a href="/contests/abc100/tasks/abc100_a">Problem A</a>`,
			reqErr: nil,
			want:   []model.ProblemId{"a"},
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
			mockRequester := testutil.MockRequester{
				ResponseStatusCode: tt.resStatusCode,
				ResponseBody:       tt.html,
				ResponseErr:        tt.reqErr,
			}
			atcoder := atcoder.NewAtcoder(&mockRequester)

			got, err := atcoder.GetProblemIds("abc100")

			if (err != nil) != tt.wantErr {
				t.Errorf("GetProblemIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProblemIds() = %v, want %v", got, tt.want)
				return
			}

			if !reflect.DeepEqual(mockRequester.CalledWith.URL.String(), "https://atcoder.jp/contests/abc100/tasks") {
				t.Errorf("GetProblemIds() = %v, want %v", mockRequester.CalledWith.URL.String(), "https://atcoder.jp/contests/abc100/tasks")
				return
			}
		})
	}

}

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
			mockRequester := testutil.MockRequester{
				ResponseStatusCode: tt.resStatusCode,
				ResponseErr:        tt.reqErr,
			}
			atcoder := atcoder.NewAtcoder(&mockRequester)

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
