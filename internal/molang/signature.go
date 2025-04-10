package molang

import (
	"log"
	"regexp"
	"strings"

	"github.com/ink0rr/rockide/internal/sliceutil"
)

type Signature string

type Parameter struct {
	Label string
	Type  string
}

var signaturePattern = regexp.MustCompile(`^\(|\):.*$`)

func (s Signature) GetParams() []Parameter {
	return sliceutil.Map(
		strings.Split(signaturePattern.ReplaceAllString(string(s), ""), ", "),
		func(s string) Parameter {
			label, paramType, ok := strings.Cut(strings.Replace(s, "[]", "", -1), ": ")
			if !ok {
				log.Panicf("invalid molang signature: %s", s)
			}
			return Parameter{Label: label, Type: paramType}
		})
}
