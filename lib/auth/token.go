package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"bitbucket.org/novatrixbr/cryptoshortner/conf"
	"bitbucket.org/novatrixbr/cryptoshortner/lib/contx"
	"github.com/dgrijalva/jwt-go"
	"github.com/novatrixtech/cryptonx"
)

// CreateJWTCookie create cookie with jwt token
func CreateJWTCookie(jwtID string, issuer string, ctx *contx.Context) {
	ip := ctx.RemoteAddr()
	expireCookie := time.Now().Add(time.Hour * 1)
	signedToken := generateJWTToken(jwtID, ip, issuer)
	cookie := http.Cookie{Name: cookie_name, Value: signedToken, Expires: expireCookie, HttpOnly: true}
	http.SetCookie(ctx.Resp, &cookie)

}

// GenerateJWTToken generate jwt token
func GenerateJWTToken(app *App, ctx *contx.Context) string {
	ip := ctx.RemoteAddr()
	return generateJWTToken(app.Id, ip, app.Name)
}

// InvalidateJWTToken invalidate jwt token
func InvalidateJWTToken(ctx *contx.Context) {
	deleteCookie := http.Cookie{Name: cookie_name, Value: "none", Expires: time.Now()}
	http.SetCookie(ctx.Resp, &deleteCookie)
}

func generateJWTToken(jwtID string, ip string, issuer string) string {
	expireToken := time.Now().Add(time.Hour * 1).Unix()

	if issuer == "" {
		issuer = "localhost:8080"
	}
	claims := Claims{
		Ip: ip,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    issuer,
			Id:        jwtID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(conf.Cfg.Section("").Key("oauth_key").Value()))
	return signedToken
}

// ClientDecrypter decrypt client token
func ClientDecrypter(key, clientID, clientSecret string) (name, id string, err error) {
	text, err := cryptonx.Decrypter(key, clientSecret, clientID)
	if err != nil {
		return "", "", err
	}
	values := strings.Split(string(text), "|")
	name = values[0]
	id = values[1]
	return
}

//ClientEncrypter encrypts new client
func ClientEncrypter(key, appName, appID string) (clientID, clientSecret string, err error) {
	clientID, clientSecret, err = cryptonx.Encrypter(key, appName+"|"+appID)
	if err != nil {
		log.Println("[ClientEncrypter] Erro ao encriptar texto: ", err.Error())
		return
	}
	return
}

func parse(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("Unexpected Signing method")
	}
	return []byte(conf.Cfg.Section("").Key("oauth_key").Value()), nil
}
