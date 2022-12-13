package services

import (
	"content/config"
	"content/feature/user/domain"
	"errors"
	lo "log"
	"strings"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &userService{
		qry: repo,
	}
}

// Register implements domain.Service
func (us *userService) Register(newUser domain.UserCore) (domain.UserCore, error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("error on bcrypt", err.Error())
		return domain.UserCore{}, errors.New("cannot encrypt password")
	}
	newUser.Password = string(generate)
	newUser.Role = 0
	orgPass := newUser.Password
	res, err := us.qry.AddUser(newUser)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			return domain.UserCore{}, errors.New("already exist")
		}
		return domain.UserCore{}, errors.New("some problem on database")
	}
	res.Password = orgPass
	return res, nil
}

// Login implements domain.Service
func (us *userService) Login(existUser domain.UserCore) (domain.UserCore, error) {
	res, err := us.qry.GetUser(existUser)
	lo.Println("hasil res", res)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "found") {
			return domain.UserCore{}, errors.New("Failed. Email or Password not found.")
		} else if strings.Contains(err.Error(), "table") {
			return domain.UserCore{}, errors.New("Failed. Email or Password not found.")
		}
		return domain.UserCore{}, errors.New("email not found")
	} else {
		if res.ID == 0 {
			return domain.UserCore{}, errors.New("email not found")
		}
		err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(existUser.Password))
		if err != nil {
			log.Error(err, " wrong password")
			return domain.UserCore{}, errors.New("wrong password")
		}
		return res, nil
	}
}

// MyProfile implements domain.Service
func (us *userService) MyProfile(userID uint) (domain.UserCore, error) {
	res, err := us.qry.GetMyUser(userID)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return domain.UserCore{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.UserCore{}, errors.New("no data")
		}
	}
	return res, nil
}

// AllUser implements domain.Service
func (us *userService) AllUser(email string) ([]domain.UserCore, error) {
	res, err := us.qry.GetAllUser(email)
	if err == gorm.ErrRecordNotFound {
		log.Error(err.Error())
		return nil, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Error(err.Error())
		return nil, errors.New(config.DATABASE_ERROR)
	}

	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New(config.DATA_NOTFOUND)
	}

	return res, nil
}

func (us *userService) MyPoint(userID uint) (domain.UserCore, error) {
	res, err := us.qry.GetMyPoint(userID)
	if err != nil {
		if strings.Contains(err.Error(), "table") {
			return domain.UserCore{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.UserCore{}, errors.New("no data")
		}
	}
	return res, nil
}
