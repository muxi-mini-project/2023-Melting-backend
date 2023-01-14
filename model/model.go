package model

import (
	"main/model/db"
)

// GetSth is used to get something from database.
// The "something" must fulfil interface "sth", which has method "TableName" and "GetKey".
func GetSth[T sth](value T) T {
	pk, id := value.GetKey()
	db.DB.Find(&value, pk+" = ?", id)
	return value
}

func UpdateSth[T sth](value T) error {
	pk, id := value.GetKey()
	result := db.DB.Table(value.TableName()).Updates(value).Where(pk+" = ?", id)
	return result.Error
}

func CreateSth[T sth](value T) error {
	var x = new(T) // <- Used to fix nil pointer panic.
	*x = value     //	Don't touch it.
	result := db.DB.Table(value.TableName()).Create(x)
	return result.Error
}

func DeleteSth[T sth](value T) error {
	result := db.DB.Table(value.TableName()).Delete(&value)
	return result.Error
}

func GetManySth[T sth](value T) ([]T, int) {
	pk, id := value.GetKey()
	data := make([]T, 100)
	result := db.DB.Table(value.TableName()).Where(pk+" = ?", id).Scan(data)
	return data, int(result.RowsAffected)
}

func GetProposals(uid int) ([]db.ProposalInfo, int) {
	data := make([]db.ProposalInfo, 100)
	result := db.DB.Table(db.TableNameProposalInfo).Where("uid = ?", uid).Scan(data)
	return data, int(result.RowsAffected)
}

func GetFromUsers(value db.User) db.User {
	db.DB.Table(db.TableNameUser).Find(&value, "nick_name = ?", value.NickName)
	return value
}

func GetGames(value db.Game) []db.Game {
	data := make([]db.Game, 100)
	db.DB.Table(value.TableName()).Where(
		"venue = ?", value.Venue,
	).Find(&data,
		"time LIKE ? or crowd LIKE ?",
		value.Time, value.Crowd,
	).Limit(99)
	return data
}
