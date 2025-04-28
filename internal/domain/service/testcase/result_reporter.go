package testcase

import (
	"bytes"
	"fmt"

	"github.com/ppdx999/kyopro/internal/domain/model"
)

type resultReporter struct{}

func (r *resultReporter) ReportTestCaseResult(
	result []byte,
	t *model.TestCase,
) string {
	buf := bytes.NewBuffer(nil)
	if bytes.Equal(result, t.Want) {
		fmt.Fprintf(buf, "✅ Test %s passed\n", t.ID)
	} else {
		fmt.Fprintf(buf, "❌ Test %s failed\n", t.ID)
		fmt.Fprintln(buf, "Input:")
		fmt.Fprintf(buf, "%q\n", t.Input)
		fmt.Fprintln(buf, "Want:")
		fmt.Fprintf(buf, "%q\n", t.Want)
		fmt.Fprintln(buf, "Got:")
		fmt.Fprintf(buf, "%q\n", result)
	}
	return buf.String()
}
