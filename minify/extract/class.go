package extract

import (
	"os"

	"golang.org/x/net/html"
)

func Class(file *os.File) map[string]struct{} {
	classesMap := make(map[string]struct{})
	htmlTokens := html.NewTokenizer(file)
loop:
	for {
		tt := htmlTokens.Next()
		switch tt {
		case html.ErrorToken:
			break loop
		case html.StartTagToken:
			t := htmlTokens.Token()
			putClassInMap(&t, classesMap)
		}
	}
	return classesMap
}
