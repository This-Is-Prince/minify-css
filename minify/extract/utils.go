package extract

import "strings"

func foundBodyStart(line *string) {
	if !f.isBodyFound && isFound(line, "<body") {
		f.isBodyFound = true
	}
}

func foundBodyEnd(line *string) {
	if f.isBodyFound && isFound(line, "</body>") {
		f.isBodyFound = false
	}
}

func isFound(line *string, delim string) bool {
	return strings.Contains(*line, delim)
}
