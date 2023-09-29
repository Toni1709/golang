package configs

import "belajar_golang/models/user"

func initMigrate() {
	DB.AutoMigrate(&user.User{})
}
