package main

import (
	"fmt"
	"main/service"
)

func main() {
	t, _ := service.Newtoken(16)
	token, claims, _ := service.Parsetoken(t)
	fmt.Println(token.Valid)
	fmt.Println(claims)
}
