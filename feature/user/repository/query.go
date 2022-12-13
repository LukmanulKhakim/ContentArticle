package repository

import (
	"content/feature/user/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.Repository {
	return &repoQuery{
		db: db,
	}
}

// AddUser implements domain.Repository
func (rq *repoQuery) AddUser(newUser domain.UserCore) (domain.UserCore, error) {
	var cnv User = FromDomain(newUser)
	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("error on adding user", err.Error())
		return domain.UserCore{}, err
	}
	return ToDomain(cnv), nil
}

// GetUser implements domain.Repository
func (rq *repoQuery) GetUser(existUser domain.UserCore) (domain.UserCore, error) {
	var cnv User
	if err := rq.db.Table("users").First(&cnv, "email = ?", existUser.Email).Error; err != nil {
		log.Error("error on get user login", err.Error())
		return domain.UserCore{}, nil
	}
	return ToDomain(cnv), nil
}

// GetMyUser implements domain.Repository
func (rq *repoQuery) GetMyUser(userID uint) (domain.UserCore, error) {
	var resQuery User
	if err := rq.db.First(&resQuery, "id = ?", userID).Error; err != nil {
		log.Error("error on get my user", err.Error())
		return domain.UserCore{}, err
	}
	return ToDomain(resQuery), nil
}

// GetAllUser implements domain.Repository
func (rq *repoQuery) GetAllUser(email string) ([]domain.UserCore, error) {
	var resQry []User
	if err := rq.db.Table("users").Select("users.id", "users.fullname", "users.email", "users.point").Scan(&resQry).Where("role = 0 AND email LIKE ?", "%"+email+"%").Find(&resQry).Error; err != nil {
		log.Error("Error on All user", err.Error())
		return nil, err
	}
	return ToDomainArray(resQry), nil
}

// GetMyPoint implements domain.Repository
func (rq *repoQuery) GetMyPoint(userID uint) (domain.UserCore, error) {
	var resQry User
	if err := rq.db.Table("contents").Where("user_id = ? ", userID).Select("SUM(point_art)").Scan(&resQry).Error; err != nil {
		log.Error("Error on All user", err.Error())
		return domain.UserCore{}, err
	}
	return ToDomain(resQry), nil
}
