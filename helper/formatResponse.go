package helper

import "regexp"

func FormatResponse(response string) string {
	re := regexp.MustCompile("^`|`$")
	return re.ReplaceAllString(response, "")
}