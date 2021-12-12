package routes

import "regexp"

var (
	DetectNgrok = regexp.MustCompile(`(https:)([/|.|\w|\s|-])*\.(?:io)`) // this is the regex for get the url
)

type url struct {
	Url   string
	Urlos string
	Urlip string
}
