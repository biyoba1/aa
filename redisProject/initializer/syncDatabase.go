package initializer

import "github.com/biyoba1/redisProject/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Person{})
}
