package rule

import (
	"strings"
	"unicode"

	"github.com/astrocorp42/flint/lint"
)

type NoWhitespaces struct{}

func (r NoWhitespaces) Apply(file lint.File) []lint.Issue {
	issues := []lint.Issue{}

	if strings.IndexFunc(file.Name, unicode.IsSpace) != -1 {
		issue := lint.Issue{
			File:         file,
			RuleName:     r.Name(),
			Explaination: "Unexpected whitespace in name",
		}
		issues = append(issues, issue)
	}

	return issues
}

func (_ NoWhitespaces) Name() string {
	return "no_whitespaces"
}
