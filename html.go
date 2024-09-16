package toolkits

import "regexp"

// RemoveHTMLTags removes all HTML tags from a string
func RemoveHTMLTags(input string) string {
	re := regexp.MustCompile("<[^>]*>")
	result := re.ReplaceAllString(input, "")
	return result
}
