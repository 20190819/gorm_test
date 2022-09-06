package main

import (
	"gorm_example/models"
	"time"
)

func main() {
	user := models.User{
		Name:     "lisi",
		Birthday: time.Now(),
	}
	user.AddUser()
}
