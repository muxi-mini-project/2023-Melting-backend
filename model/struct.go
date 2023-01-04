package model

type QNconfig struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Domain    string
}

type sth interface {
	TableName() string
	getKey() (string, int)
}
