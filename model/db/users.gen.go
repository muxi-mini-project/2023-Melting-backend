// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	UID         int32  `gorm:"column:uid;primaryKey;autoIncrement:true" json:"uid"` // 序号
	Photo       string `gorm:"column:photo" json:"photo"`                           // 头像
	NickName    string `gorm:"column:nick_name;default:昵称" json:"nick_name"`        // 最多七个汉字
	Position    string `gorm:"column:position;default:职位" json:"position"`          // 职位
	Tag         string `gorm:"column:tag" json:"tag"`                               // 标签
	Auth        string `gorm:"column:auth" json:"auth"`
	Muxipass    int32  `gorm:"column:muxipass" json:"muxipass"`
	Studentid   int32  `gorm:"column:studentid" json:"studentid"`
	Qq          string `gorm:"column:qq" json:"qq"`
	Description string `gorm:"column:description;not null" json:"description"`
}

// TableName User's table name
func (User) TableName() string {
	return TableNameUser
}
