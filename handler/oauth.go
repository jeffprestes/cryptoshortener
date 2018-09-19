package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"bitbucket.org/novatrixbr/cryptoshortner/conf"
	"bitbucket.org/novatrixbr/cryptoshortner/lib/auth"
	"bitbucket.org/novatrixbr/cryptoshortner/lib/contx"
)

//Oauth generate client access token
func Oauth(ctx *contx.Context) {
	key := conf.Cfg.Section("").Key("oauth_key").Value()
	id, secret, _ := ctx.Req.BasicAuth()
	credentials := auth.Oauth{
		Id:     id,
		Secret: secret,
	}
	appName, appID, _ := auth.ClientDecrypter(key, id, secret)
	data := auth.DB[credentials]
	//log.Println("id: ", id, " secret: ", secret, " appName: ", appName, " appID: ", appID, " - Data: ", data)
	if data.Id == appID && data.Name == appName {
		token := auth.GenerateJWTToken(data, ctx)
		ctx.JSON(http.StatusOK, map[string]string{"access_token": token})
		return
	}
	ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid Credentials"})
	return

}

//NewAppClient generate client access token
func NewAppClient(ctx *contx.Context) {
	body, err := ctx.Req.Body().Bytes()
	defer ctx.Req.Body().ReadCloser()
	if err != nil {
		log.Println("[NewAppClient] Erro ao pegar o conteudo do body: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	app := auth.App{}
	if err := json.Unmarshal(body, &app); err != nil {
		log.Println("[NewAppClient] Error Parsing ", err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	key := conf.Cfg.Section("").Key("oauth_key").Value()
	id, secret, _ := auth.ClientEncrypter(key, app.Name, app.Id)
	credentials := auth.Oauth{
		Id:     id,
		Secret: secret,
	}
	ctx.JSON(http.StatusOK, credentials)
}
