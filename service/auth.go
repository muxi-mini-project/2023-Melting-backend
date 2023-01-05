package service

import (
	"main/model"
)

func Login(data model.Auth)(string,error) {
	s := data.Auth
	loginAuth := model.Auth{
		UID: data.UID,
	}
	model.GetSth(&loginAuth)
	if s == loginAuth.Auth {
		t, err := Newtoken(int(loginAuth.UID))
		if err != nil {
			return "",err
		} else {
			return t,nil
		}
	} else {
		return "",model.ErrAuthIncorrect
	}
}
