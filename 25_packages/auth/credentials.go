package auth

import "fmt"

func Login_with_credentials(username string, password string){
	fmt.Println("login user using", username, password)
}

// this func can be used anywhere, where package auth is used


//  1st alphabet in function name go convention
/*
	small 1st aplhabet :------>  local scope in package (only work in current package)
	capital            :------>  global scope in packages (can work anywhere, where this package is called)
*/