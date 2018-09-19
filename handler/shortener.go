package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2"

	"bitbucket.org/novatrixbr/cryptoshortner/lib/contx"
	"bitbucket.org/novatrixbr/cryptoshortner/model"
	"bitbucket.org/novatrixbr/cryptoshortner/repo"
)

//GetAccount gets crypto account from sequence informed
func GetAccount(ctx *contx.Context) {
	nick := strings.TrimSpace(ctx.ParamsEscape(":nickname"))
	if len(nick) < 2 {
		ctx.RawData(http.StatusBadRequest, []byte("Invalid sequence"))
		return
	}

	var network string
	testnet := ctx.QueryBool("test")
	if !testnet {
		network = "mainnet"
	} else {
		testnet := ctx.QueryTrim("network")
		if testnet == "rinkeby" {
			network = "rinkeby"
		}
	}

	if network == "" {
		ctx.RawData(http.StatusBadRequest, []byte("Invalid sequence"))
		return
	}

	account, err := repo.QueryNicknamesInEthereum(nick, network)
	if err != nil {
		if err == mgo.ErrNotFound {
			ctx.RawData(http.StatusNotFound, []byte("Not found"))
			return
		}
		ctx.RawData(http.StatusInternalServerError, []byte("Internal error. Please contact webmaster"))
		return
	}

	if !strings.HasPrefix(account.AccountID, "0x") {
		account.Network = "Bitcoin"
	}

	exhibitionMode := ctx.QueryTrim("v")
	if exhibitionMode == "json" {
		ctx.JSON(http.StatusOK, account)
		return
	} else if exhibitionMode == "momcode" {
		if (len(nick[1:]) % 2) != 0 {
			ctx.RawData(http.StatusBadRequest, []byte("Invalid sequence"))
			return
		}
		//TODO: create a page to show the visualization in creation mode
		ctx.JSON(http.StatusOK, account)
		return
	}

	if strings.ToLower(account.Network) == "ethereum" {
		var ethscanPrefix string
		if network == "rinkeby" {
			ethscanPrefix = "rinkeby."
		}
		ctx.Redirect("https://"+ethscanPrefix+"etherscan.io/address/"+account.AccountID, http.StatusFound)
		return
	} else if strings.ToLower(account.Network) == "bitcoin" {
		ctx.Redirect("https://www.blockchain.com/btc/address/"+account.AccountID, http.StatusFound)
		return
	}
}

//InsertAccount inserts a new nickname to an account
func InsertAccount(ctx *contx.Context) {
	body, err := ctx.Req.Body().Bytes()
	defer ctx.Req.Body().ReadCloser()
	if err != nil {
		log.Println("[InsertAccount] Erro ao pegar o conteudo do body: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	account := model.Account{}
	if err := json.Unmarshal(body, &account); err != nil {
		log.Println("[InsertAccount] Error Parsing ", err.Error())
		if strings.Contains(err.Error(), "duplicate key error") {
			ctx.RawData(http.StatusBadRequest, []byte("duplicated key"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	err = repo.InsertNickname(account)
	if err != nil {
		log.Println("[InsertAccount] Erro ao inserir um novo nickname: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.RawData(http.StatusOK, []byte("ok"))
}
