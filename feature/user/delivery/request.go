package delivery

import "content/feature/user/domain"

type UserFormat struct {
	Fullname string `json:"fullname" form:"fullname"  `
	Email    string `json:"email" form:"email" `
	Password string `json:"password" form:"password"`
}

type LoginFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToDomain(i interface{}) domain.UserCore {
	switch i.(type) {
	case UserFormat:
		cnv := i.(UserFormat)
		return domain.UserCore{Fullname: cnv.Fullname, Email: cnv.Email, Password: cnv.Password}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.UserCore{Email: cnv.Email, Password: cnv.Password}
	}
	return domain.UserCore{}
}
