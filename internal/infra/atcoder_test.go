package infra_test

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/ppdx999/kyopro/internal/infra"
	"github.com/ppdx999/kyopro/internal/infra/testutil"
	"github.com/ppdx999/kyopro/internal/model"
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
			atcoder := infra.NewAtcoder(&testutil.MockRequester{})

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
			atcoder := infra.NewAtcoder(&mockRequester)

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
