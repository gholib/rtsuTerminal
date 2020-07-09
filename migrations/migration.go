package migrations

import "github.com/gholib/rtsuTerminal/models"

func Automigrate()  {
db := models.Connect()
defer db.Close()
db.DropTableIfExists(&models.Terminal{}, &models.Service{})
db.AutoMigrate(&models.Terminal{}, &models.Service{})

	
}

//
//func Run() {
//
//
//	database.DB.AutoMigrate(models.User{})
//}
