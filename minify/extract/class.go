package extract

import (
	"os"
	"strings"

	"github.com/This-Is-Prince/minify-css/minify/utils"
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
			for _, v := range t.Attr {
				if v.Key == "class" {
					classArr := strings.Split(v.Val, " ")
					for _, class := range classArr {
						class = strings.TrimSpace(class)
						if class != "" {
							classesMap[class] = utils.Exists
						}
					}
				}
			}
		}
	}
	return classesMap
}
