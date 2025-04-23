package atcoder_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ppdx999/kyopro/internal/atcoder"
	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/testutil"
)

func TestGetTestCases(t *testing.T) {
	tests := []struct {
		name          string
		html          string
		reqErr        error
		resStatusCode int
		want          []*model.TestCase
		wantErr       bool
	}{
		{
			name: "正常系: １ケース",
			html: `
				<h3>入力例 1</h3>
				<pre>1 2</pre>
				<h3>出力例 1</h3>
				<pre>3</pre>
			`,
			resStatusCode: 200,
			want: []*model.TestCase{
				{
					ID:    "1",
					Input: []byte("1 2"),
					Want:  []byte("3"),
				},
			},
		},
		{
			name: "正常系: 複数ケース 空白を含む",
			html: `
				<h3>入力例 1</h3> <pre> 10 20 </pre>
				<h3>出力例 1</h3>   <pre>30
</pre>
				<h3>入力例 2</h3><pre> -5 5 </pre>
				<h3>出力例 2</h3><pre>0</pre>
			`,
			resStatusCode: 200,
			want: []*model.TestCase{
				{
					ID:    "1",
					Input: []byte("10 20"),
					Want:  []byte("30"),
				},
				{
					ID:    "2",
					Input: []byte("-5 5"),
					Want:  []byte("0"),
				},
			},
		},
		{
			name:    "異常系: 入力のみ",
			html:    `<h3>入力例 1</h3> <pre> 10 20 </pre>`,
			wantErr: true,
		},
		{
			name: "準異常系: 入力出力ともになし",
			html: ``,
			want: nil,
		},
		{
			name:    "request エラー",
			reqErr:  errors.New("request error"),
			wantErr: true,
		},
		{
			name:          "ステータスコードエラー",
			resStatusCode: 500,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockReq := &testutil.MockRequester{
				ResponseStatusCode: tt.resStatusCode,
				ResponseBody:       tt.html,
				ResponseErr:        tt.reqErr,
			}
			atcoder := atcoder.NewAtcoder(mockReq)

			got, err := atcoder.GetTestCases("", "")
			if (err != nil) != tt.wantErr {
				t.Errorf("Atcoder.GetTestCases() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Atcoder.GetTestCases() = %v, want %v", got, tt.want)
			}
		})
	}
}
