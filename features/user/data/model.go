package data

import "altaproject2/domain"

type User struct {
	ID           int    `json:"id" form:"id" gorm:"primaryKey:autoIncrement"`
	Fullname     string `json:"fullname" form:"fullname" validate:"required"`
	Username     string `json:"username" form:"username" gorm:"unique"`
	Email        string `json:"email" form:"email" validate:"required" gorm:"unique"`
	Address      string `json:"address" form:"address" validate:"required"`
	PhotoProfile string `json:"photoprofile" form:"photoprofile"`
	Password     string `json:"password" form:"password" validate:"required"`
	Role         string
	// MyOrder []data.Order
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:           u.ID,
		Fullname:     u.Fullname,
		Username:     u.Username,
		Email:        u.Email,
		Address:      u.Address,
		PhotoProfile: u.PhotoProfile,
		Password:     u.Password,
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
	res.ID = data.ID
	res.Fullname = data.Fullname
	res.Username = data.Username
	res.Email = data.Email
	res.Address = data.Address
	res.PhotoProfile = data.PhotoProfile
	res.Password = data.Password
	return res
}
