package regex

import (
	"regexp"
)

var (
	EmailRegex *regexp.Regexp
)

func init(){
	EmailRegex, _ = regexp.Compile("[^@]+@[^\\.]+\\..+")
}