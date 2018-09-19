package repo

import (
	"log"

	"github.com/ethereum/go-ethereum/common"

	"bitbucket.org/novatrixbr/cryptoshortner/contract"

	"bitbucket.org/novatrixbr/cryptoshortner/conf"
	"bitbucket.org/novatrixbr/cryptoshortner/model"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/mgo.v2/bson"
)

//QueryNicknamesInEthereum gets crypto account from repo based on nickname informed stored at Ethereum network
func QueryNicknamesInEthereum(nick string, network string) (account model.Account, err error) {
	err = nil
	client, err := ethclient.Dial("https://" + network + ".infura.io/QPF0qjGpH9OjFuuMrCse")
	if err != nil {
		log.Printf("[QueryNicknamesInEthereum] Houve falha ao conectar com Rinkeby via infuria: %+v", err)
		return
	}
	var enderecoContrato common.Address
	if network == "rinkeby" {
		enderecoContrato = common.HexToAddress("0x4ead784006cb920965cb1b45a9f6923330a3ccf0")
	} else {
		enderecoContrato = common.HexToAddress("0x023fa9e2a97799b3D87B3fa35674b50B8b5c9f4E")
	}
	contrato, err := contract.NewJNS(enderecoContrato, client)
	if err != nil {
		log.Printf("[QueryNicknamesInEthereum] Houve falha ao conectar com o contrato na rede Rinkeby via infuria: %+v", err)
		return
	}
	temp, err := contrato.GetAddress(nil, nick)
	if err != nil {
		log.Printf("[QueryNicknamesInEthereum] Houve falha ao obter o endereco com o nick informado: %+v", err)
		return
	}
	account.AccountID = temp.String()
	account.Network = "Ethereum"
	account.Nickname = nick
	return
}

//QueryNicknames gets crypto account from repo based on nickname informed
func QueryNicknames(nick string) (account model.Account, err error) {
	col, err := conf.GetMongoCollection("listaccounts")
	if err != nil {
		log.Printf("[QueryNicknames] Erro ao obter a colecao: %s\n", err.Error())
		return
	}
	defer col.Database.Session.Close()

	err = col.Find(bson.M{"nickname": nick}).One(&account)
	if err != nil {
		log.Printf("[QueryNicknames] Erro ao fazer a busca e o bind da colecao: %s\n", err.Error())
		return
	}
	return
}

//InsertNickname inserts a new account nickname into collection
func InsertNickname(account model.Account) (err error) {
	err = nil
	col, err := conf.GetMongoCollection("listaccounts")
	if err != nil {
		log.Printf("[InsertNickname] Erro ao obter a colecao: %s\n", err.Error())
		return
	}
	defer col.Database.Session.Close()

	err = col.Insert(account)
	if err != nil {
		log.Printf("[InsertNickname] Erro ao inserir um novo nickname colecao: %s\n", err.Error())
		return
	}
	return

}
