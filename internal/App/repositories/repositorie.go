package repositories

import (
	"auth/internal/App/models"
	"auth/internal/utils"
)

func InitDatabase() {
	dbName := "auth.db"
	gorm := utils.OrmInit(dbName)
	//gorm.AutoMigrate(models.User{})

	u1 := models.NewUser("mouhamed", "syllamouhamed99@gmail.com", "1234")

	gorm.Insert(u1)
}
