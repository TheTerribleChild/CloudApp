package regexutil

import (
	"regexp"
)

var (
	EmailRegex *regexp.Regexp
)

func init(){
	EmailRegex, _ = regexp.Compile("[^@]+@[^\\.]+\\..+")
}