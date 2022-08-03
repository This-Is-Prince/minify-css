package extract

import (
	"strings"

	"github.com/This-Is-Prince/minify-css/minify/utils"
	"golang.org/x/net/html"
)

func putClassInMap(t *html.Token, classesMap map[string]struct{}) {
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
