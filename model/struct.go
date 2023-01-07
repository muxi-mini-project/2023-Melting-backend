package model

import "main/model/db"

type QNconfig struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Domain    string
}

type sth interface {
	db.User | db.Template | db.ProposalInfo | db.Tag | db.Question
	TableName() string
	GetKey() (string, int)
}
