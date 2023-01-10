package db

func (data User) GetKey() (string, int) {
	return "uid", int(data.UID)
}

func (data ProposalInfo) GetKey() (string, int) {
	return "info_id", int(data.InfoID)
}

func (data Tag) GetKey() (string, int) {
	return "tid", int(data.Tid)
}

func (data Question) GetKey() (string, int) {
	return "qid", int(data.Qid)
}

func (data Template) GetKey() (string, int) {
	return "temid", int(data.Temid)
}

func (data Game) GetKey() (string, int) {
	return "gameid", int(data.Gameid)
}