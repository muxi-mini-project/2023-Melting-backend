package model

func (data *User) getKey() (string, int) {
	return "uid", int(data.UID)
}
func (data *ProposalArr) getKey() (string, int) {
	return "arr_id", int(data.ArrID)
}
func (data *ProposalInfo) getKey() (string, int) {
	return "info_id", int(data.InfoID)
}
func (data *Tag) getKey() (string, int) {
	return "tid", int(data.Tid)
}
func (data *Question) getKey() (string, int) {
	return "qid", int(data.Qid)
}
