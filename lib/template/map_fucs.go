package template

import (
	"html/template"

	"bitbucket.org/novatrixbr/cryptoshortner/lib/contx"
)

// FuncMaps to view
func FuncMaps() []template.FuncMap {
	return []template.FuncMap{
		map[string]interface{}{
			"Tr": contx.I18n,
		}}
}
