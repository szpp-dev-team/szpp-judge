package judge

import "strings"

type Checker struct {
	Output       string
	ExpectOutput string
}

// TODO: testlib を使う
func (c *Checker) JudgeNormal(caseInsensitive bool) bool {
	if caseInsensitive {
		return strings.EqualFold(c.Output, c.ExpectOutput)
	}
	return c.Output == c.ExpectOutput
}
