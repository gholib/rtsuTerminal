package migrations

import "github.com/gholib/rtsuTerminal/models"

func Automigrate()  {
db := models.Connect()
defer db.Close()
db.AutoMigrate(&models.Terminal{})

	
}

//
//func Run() {
//
//
//	database.DB.AutoMigrate(models.User{})
//}
