package htmlutil

import (
    "html/template"
)

func Unescape(htmlStr string) template.HTML {
    return template.HTML(htmlStr)
}
