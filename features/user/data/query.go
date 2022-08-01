package data

import (
	"fmt"
	"log"

	"altaproject2/domain"
	"altaproject2/features/common"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserData {
	return &userData{
		db: db,
	}
}

func (ud *userData) Login(email string, password string) (username string, role string, token string, err error) {
	userData := User{}
	ud.db.Where("email = ?", email).First(&userData)
	check := common.CheckPasswordHash(password, userData.Password)

	if !check {
		return "", "", "", fmt.Errorf("error")
	}

	token = common.GenerateToken(domain.User(userData))

	return userData.Username, userData.Role, token, nil
}

func (ud *userData) Delete(userID int) bool {
	err := ud.db.Where("ID = ?", userID).Delete(&User{})
	if err.Error != nil {
		log.Println("cannot delete content", err.Error.Error())
		return false
	}
	if err.RowsAffected < 1 {
		log.Println("No content deleted", err.Error.Error())
		return false
	}

	return true
}
