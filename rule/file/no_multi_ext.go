package file

import (
	"strings"

	"github.com/z0mbie42/flint/lint"
)

type NoMultiExt struct{}

func (r NoMultiExt) Apply(file lint.File) []lint.Issue {
	dotCount := strings.Count(file.Name, ".")
	issues := []lint.Issue{}

	if file.IsDir {
		return issues
	}

	if dotCount > 1 {
		issue := lint.Issue{
			File:         file,
			RuleName:     r.Name(),
			Explaination: "should not have multiple extensions (multiple . in name)",
		}
		issues = append(issues, issue)
	}

	return issues
}

func (_ NoMultiExt) Name() string {
	return "file/no_multi_ext"
}