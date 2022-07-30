package minify

import (
	"fmt"

	"github.com/This-Is-Prince/minify-css/minify/extract"
	"github.com/This-Is-Prince/minify-css/minify/htmlFile"
)

func Run() {
	file, err := htmlFile.Open()
	if err != nil {
		fmt.Println(err)
	} else {
		extract.Class(file)
	}
}
