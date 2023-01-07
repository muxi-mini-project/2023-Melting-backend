package model

import (
	"main/model/db"
)

// TODO Not working
// GetSth is used to get something from database.
// The "something" must fulfil interface "sth", which has method "TableName" and "GetKey".
func GetSth[T sth](value T) T {
	pk, id := value.GetKey()
	db.DB.Find(&value, pk+" = ?", id)
	return value
}

func GetFromUsers(value db.User) db.User {
	db.DB.Table(db.TableNameUser).Find(&value, "nick_name = ?", value.NickName)
	return value
}

func UpdateSth[T sth](value T) {
	pk, id := value.GetKey()
	db.DB.Table(value.TableName()).Updates(value).Where(pk+" = ?", id)
}

func CreateSth[T sth](value T) {
	var x = new(T) // <- Used to fix nil pointer panic.
	*x = value     //	Don't touch it.
	db.DB.Table(value.TableName()).Create(x)
}

func DeleteSth[T sth](value T) {
	db.DB.Table(value.TableName()).Delete(&value)
}

func GetProposals(uid int) ([]db.ProposalInfo, int) {
	data := make([]db.ProposalInfo, 100)
	result := db.DB.Table(db.TableNameProposalInfo).Where("uid = ?", uid).Scan(data)
	return data, int(result.RowsAffected)
}
