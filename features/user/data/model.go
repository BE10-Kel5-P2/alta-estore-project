package data

import (
	"altaproject2/domain"
	cartdata "altaproject2/features/cart/data"
	"altaproject2/features/order/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `json:"username" form:"username" gorm:"unique"`
	Email        string `json:"email" form:"email" validate:"required" gorm:"unique"`
	Address      string `json:"address" form:"address" validate:"required"`
	PhotoProfile string `json:"photoprofile" form:"photoprofile"`
	Password     string `json:"password" form:"password" validate:"required"`
	Role         string
	Cart         []cartdata.Cart `gorm:"foreignKey:Userid"`
	MyOrder      []data.Order    `gorm:"foreignKey:Userid"`
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:           int(u.ID),
		Username:     u.Username,
		Email:        u.Email,
		Address:      u.Address,
		PhotoProfile: u.PhotoProfile,
		Password:     u.Password,
		Role:         u.Role,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.User) User {
	var res User
	res.ID = uint(data.ID)
	res.Username = data.Username
	res.Email = data.Email
	res.Address = data.Address
	res.PhotoProfile = data.PhotoProfile
	res.Password = data.Password
	res.Role = data.Role
	return res
}
