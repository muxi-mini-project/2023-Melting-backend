package model

import (
	"fmt"
	"path"
	"runtime"
)

type err struct {
	Code int
	Typ  string
}

func (e err) Error() string {
	return e.Typ
}

func ErrorSender() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func Debugger(skip int) (fileName string, funcName string, line int) {
	pc, file, line, _ := runtime.Caller(skip)
	fmt.Println(pc, file)
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	return fileName, funcName, line
}

var ErrAuthIncorrect = err{
	Code: 403,
	Typ:  "ErrAuthIncorrect",
}
var ErrAuthInvalid = err{
	Code: 403,
	Typ:  "ErrAuthInvalid",
}
var ErrBadRequest = err{
	Code: 400,
	Typ:  "ErrBadRequest",
}
var ErrNameOccupied = err{
	Code: 403,
	Typ:  "Username has been occupied",
}
var ErrNotAuthorized = err{
	Code: 403,
	Typ:  "Operation requires authentication",
}
