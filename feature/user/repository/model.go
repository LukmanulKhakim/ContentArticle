package repository

import (
	"content/feature/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"unique"`
	Password string
	Role     uint
	Point    int
	Token    string
}

func FromDomain(du domain.UserCore) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Fullname: du.Fullname,
		Email:    du.Email,
		Password: du.Password,
		Role:     du.Role,
		Point:    du.Point,
		Token:    du.Token,
	}
}

func ToDomain(u User) domain.UserCore {
	return domain.UserCore{
		ID:       u.ID,
		Fullname: u.Fullname,
		Email:    u.Email,
		Password: u.Password,
		Point:    u.Point,
		Role:     u.Role,
		Token:    u.Token,
	}
}

func ToDomainArray(au []User) []domain.UserCore {
	var res []domain.UserCore
	for _, val := range au {
		res = append(res, domain.UserCore{
			ID:       val.ID,
			Fullname: val.Fullname,
			Email:    val.Email,
			Point:    val.Point,
			Role:     val.Role,
		})
	}
	return res
}
