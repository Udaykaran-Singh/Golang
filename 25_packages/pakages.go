package main

import (
	"fmt"
	"uday_module/25_packages/auth"
	"uday_module/25_packages/user"
)

func main() {
	// how to create package :-

	// in terminal write :-
	/* go mod init module_name   (mainly git repo given)*/

	// this made the go.mod file in coder

	auth.Login_with_credentials("Udaykaran Singh", "jddnflan@adf")

	fmt.Println(auth.GetSession())

	a := user.User{
		Email: "fkjn@kdj.com",
		Name: "uday",

	}

	fmt.Println(a)

}
