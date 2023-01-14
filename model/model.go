package model

import (
	"github.com/jinzhu/gorm"
)

//用户
type User struct {
	gorm.Model
	Id string `json:"id" binding:"required" gorm:"column:id;not null"`
	PassWord  string `json:"password" binding:"required" gorm:"column:password;not null"`
	Nick_name string `json:"nick_Name" binding:"required" gorm:"column:nick_name;not null"`
	Position string `json:"position" binding:"required" gorm:"column:position;not null"
}
//项目	
type project struct {
    gorm.Model
    User_id int `json:"user_id" binding:"required" gorm:"column:user_id;not null"`
    Proposal_id int `json:"proposal_id" binding:"required" gorm:"column:proposal_id;not null"`
    Place string `json:"place" binding:"required" gorm:"column:place;not null"`
    Time datetime `json:"time" binding:"required" gorm:"column:time;not null"`
    Flock int `json:"flock" binding:"required" gorm:"column:flock;not null"`
    Department_name string `json:"department_name" binding:"required" gorm:"column:department_name;not null"`
    Proposal_name string `json:"proposal_name" binding:"required" gorm:"column:proposal_name;not null"`
    Proposal_aim string `json:"proposal_aim" binding:"required" gorm:"column:proposal_aim;not null"`
    Inform_txt string `json:"inform_txt" binding:"required" gorm:"column:inform_txt;not null"`
    Target string `json:"target" binding:"required" gorm:"column:target;not null"`
    Rules_expression string `json:"rules_expression" binding:"required" gorm:"column:rules_expression;not null"`
    Advice&basis string `json:"advice&basis" binding:"required" gorm:"column:advice&basis;not null"`
    Optional_tables string `json:"optional_tables" binding:"required" gorm:"column:optional_tables;not null"`
}

//活动
type activity struct{
    arrangement_id int `json:"arrangement_id" binding:"required" gorm:"column:arrangement_id;not null"`
    hint_word string `json:"hint_word" binding:"required" gorm:"column:hint_word;not null"`
    game string `json:"game" binding:"required" gorm:"column:game;not null"`
    proposal_nodes string `json:"proposal_nodes" binding:"required" gorm:"column:proposal_nodes;not null"`
}

//题库
type question_warehouse {
    warehouse_id int `json:"warehouse_id" binding:"required" gorm:"column:warehouse_id;not null"`
    questions string `json:"question" binding:"required" gorm:"column:qusetion;not null"`
}


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
	result := db.DB.Table(value.TableName()).Where(pk+" = ?", id).Scan(&data)
	return data, int(result.RowsAffected)
}

func GetProposals(uid int) ([]db.ProposalInfo, int) {
	data := make([]db.ProposalInfo, 100)
	result := db.DB.Table(db.TableNameProposalInfo).Where("uid = ?", uid).Scan(&data)
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
