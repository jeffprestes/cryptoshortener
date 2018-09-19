package cors

import (
	"bitbucket.org/novatrixbr/cryptoshortner/lib/contx"
	"gopkg.in/macaron.v1"
)

func Cors() macaron.Handler {
	return func(ctx *contx.Context) {
		ctx.Resp.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Resp.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Resp.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Next()
	}
}
