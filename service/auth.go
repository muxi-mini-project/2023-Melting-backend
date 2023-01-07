package service

import (
	"main/model"
	"main/model/db"
)

// Login with uid
func LoginNative(data db.User) (string, error) {
	// TODO Login auth should be Base64 encrypted
	// s, err := B64Decode(data.Auth)
	//if err != nil {
	//	return "", err
	//}
	//ds := string(s)
	ds := data.Auth
	loginAuth := db.User{
		NickName: data.NickName,
	}
	loginAuth = model.GetFromUsers(loginAuth)
	if ds == loginAuth.Auth {
		t, err := Newtoken(int(loginAuth.UID))
		if err != nil {
			return "", err
		} else {
			return t, nil
		}
	} else {
		return "", model.ErrAuthIncorrect
	}
}

func CreateNative(data db.User) error {
	tmp := db.User{UID: 0, NickName: data.NickName}
	tmp = model.GetFromUsers(tmp)
	if tmp.UID != 0 {
		return model.ErrNameOccupied
	}
	if data.Auth == "" {
		return model.ErrAuthInvalid
	}
	model.CreateSth(data)
	return nil
}

func CreateWithQQ(data db.User) error {
	tmp := db.User{Qq: data.Qq}
	tmp = model.GetFromUsers(tmp)
	if tmp.Auth != "" {
		return model.ErrNameOccupied
	}
	if data.Auth == "" {
		return model.ErrAuthInvalid
	}
	model.CreateSth(data)
	return nil
}
