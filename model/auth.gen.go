// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAuth = "auth"

// Auth mapped from table <auth>
type Auth struct {
	UID         int32  `gorm:"column:uid;primaryKey;autoIncrement:true" json:"uid"`
	Auth        string `gorm:"column:auth" json:"auth"`
	Muxipass    int32  `gorm:"column:muxipass" json:"muxipass"`
	Studentid   int32  `gorm:"column:studentid" json:"studentid"`
	Description string `gorm:"column:description;not null" json:"description"`
}

// TableName Auth's table name
func (*Auth) TableName() string {
	return TableNameAuth
}
