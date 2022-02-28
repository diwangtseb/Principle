package main

type IUserManager interface {
	updateUser(username string, password string) // That's the wrong definition
	updateUserName(newName string)               // √ maybe should this
	updateUserEmail(newEmail string)             // √  Shooting
}

type IUserManagerPretty interface {
	updateUserName(newName string)   // √ maybe should this
	updateUserEmail(newEmail string) // √  Shooting
}

func main() {
	//# TODO
}
