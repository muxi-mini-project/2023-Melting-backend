package model

func GetUserAuth() {
	//TODO
}

func GetSth(value *sth) {
	pk, id := (*value).getKey()
	DB.Find(&value, pk+" = ?", id)
}

func UpdateSth(value sth) {
	pk, id := value.getKey()
	DB.Table(value.TableName()).Updates(value).Where(pk+" = ?", id)
}

func CreateSth(value sth) {
	DB.Table(value.TableName()).Create(value)
}

func DeleteSth(value sth) {
	DB.Table(value.TableName()).Delete(&value)
}
