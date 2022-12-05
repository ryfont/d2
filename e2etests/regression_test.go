package e2etests

import (
	"testing"
)

func testRegression(t *testing.T) {
	tcs := []testCase{
		{
			name: "dagre_id_with_newline",
			script: `
ninety\nnine
eighty\reight
seventy\r\nseven
`,
		},
	}

	runa(t, tcs)
}
