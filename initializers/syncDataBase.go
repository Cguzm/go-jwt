package initializers

import "github.com/rickyguzm/go-jwt/models"

func SyncDataBase() {
	DB.AutoMigrate(&models.User{})
}
