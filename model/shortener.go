package model

//Account represents a cryptoaccount nickname
type Account struct {
	Nickname  string `json:"nickname"`
	AccountID string `json:"account" bson:"account"`
	Network   string `json:"network"`
}
