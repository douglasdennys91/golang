package utils

import (
	"github.com/gosimple/slug"
)

func RenderSLUG(name string) string {
	text := slug.Make(name)
	return text
}