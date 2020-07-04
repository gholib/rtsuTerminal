package migrations

import (
	"github.com/gholib/rtsuTerminal/database"
	"github.com/gholib/rtsuTerminal/models"
)

func Run() {
	database.DB.AutoMigrate(models.User{})
}
