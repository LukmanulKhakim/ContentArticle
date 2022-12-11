package repository

import (
	"content/feature/article/domain"

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
	Contents []Content `gorm:"foreignKey:User_ID"`
}

type Content struct {
	gorm.Model
	Article   string
	Point_Art int
	User_ID   uint
	Fullname  string `gorm:"-:migration" gorm:"<-"`
}

func FromDomain(da domain.ContentCore) Content {
	return Content{
		Model:     gorm.Model{ID: da.ID},
		Article:   da.Article,
		Point_Art: da.Point_Art,
		User_ID:   da.User_ID,
		Fullname:  da.Fullname,
	}
}

func ToDomain(a Content) domain.ContentCore {
	return domain.ContentCore{
		ID:        a.ID,
		Article:   a.Article,
		Point_Art: a.Point_Art,
		User_ID:   a.User_ID,
		Fullname:  a.Fullname,
	}
}

func ToDomainArray(ac []Content) []domain.ContentCore {
	var res []domain.ContentCore
	for _, val := range ac {
		res = append(res, domain.ContentCore{
			ID:        val.ID,
			Article:   val.Article,
			Point_Art: val.Point_Art,
			User_ID:   val.User_ID,
			Fullname:  val.Fullname,
		})
	}
	return res
}
