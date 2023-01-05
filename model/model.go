package model


//GetSth is used to get something from database.
//The "something" must fulfilled interface "sth", which has method "TableName" and "getKey".
func GetSth(value sth) {
	pk, id := (value).getKey()
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

func GetProposals(uid int)([]ProposalInfo,int){
	data := make([]ProposalInfo,100)
	result:=DB.Table(TableNameProposalInfo).Where("uid = ?",uid).Scan(data)
	return data,int(result.RowsAffected)
}