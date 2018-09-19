package handler

import (
	"net/http"

	"bitbucket.org/novatrixbr/cryptoshortner/lib/contx"
)

//Index shows index page
func Index(ctx *contx.Context) {
	ctx.NativeHTML(http.StatusOK, "index")
}

//OpenCreatePage opens the create page where user is going to set her URL
func OpenCreatePage(ctx *contx.Context) {
	ctx.NativeHTML(http.StatusOK, "create")
}
