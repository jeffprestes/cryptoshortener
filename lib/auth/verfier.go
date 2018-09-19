package auth

import (
	"net/http"
	"strings"
	"log"

	"github.com/dgrijalva/jwt-go"
	"bitbucket.org/novatrixbr/cryptoshortner/lib/contx"
)

const cookie_name = "mercuriusAuth"

// Oauth
type Oauth struct {
	Id     string `json:"id"`
	Secret string `json:"secret"`
}

// Claims
type Claims struct {
	Ip string `json:"ip"`
	jwt.StandardClaims
}

// App
type App struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

// DB
var DB map[Oauth]*App = make(map[Oauth]*App)

// LoginRequired
func LoginRequired(ctx *contx.Context) {
	cookie, err := ctx.Req.Cookie(cookie_name)
	if err != nil {
		ctx.Redirect("/login")
		log.Println(err)
		return
	}
	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, parse)
	if err != nil {
		ctx.Redirect("/login")
		log.Println(err)
		return
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid && ctx.RemoteAddr() == claims.Ip {
		ctx.Data["jwt"] = *claims
	} else {
		ctx.Redirect("/login")
		log.Println("Cause: Invalid token")
	}
}

// LoginRequiredApi
func LoginRequiredApi(ctx *contx.Context) {
    header := ctx.Req.Header.Get("Authorization")
    if header != "" {
        splitted := strings.Split(header, " ")
		if len(splitted) < 2 {
			ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Malformed request header"})
			return
		}
		value := splitted[1]
		token, err := jwt.ParseWithClaims(value, &Claims{}, parse)
        if err != nil {
            ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
            return
        }
        if claims, ok := token.Claims.(*Claims); ok && token.Valid && ctx.RemoteAddr() == claims.Ip {
            ctx.Data["jwt"] = *claims
            return
        } else {
            ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
            return
        }
    }
    ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
}