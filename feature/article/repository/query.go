package repository

import (
	"content/feature/article/domain"

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

// Add implements domain.Repository
func (rq *repoQuery) Add(newItem domain.ContentCore) (domain.ContentCore, error) {
	var cnv Content = FromDomain(newItem)
	if err := rq.db.Table("contents").Select("article", "point_art", "user_id").Create(&cnv).Error; err != nil {
		log.Error("Error on insert user", err.Error())
		return domain.ContentCore{}, err
	}
	return ToDomain(cnv), nil
}

// Edit implements domain.Repository
func (rq *repoQuery) Edit(point domain.ContentCore, contentID uint, user uint) (domain.ContentCore, domain.User, error) {
	var cnv Content = FromDomain(point)
	var res User
	if err := rq.db.Table("contents").Where("id = ?", contentID).Updates(&cnv).Error; err != nil {
		log.Error("error on updating user", err.Error())
		return domain.ContentCore{}, domain.User{}, err
	}

	if err := rq.db.Table("contents").Where("user_id = ?", user).Select("sum(point_art)").Updates(&res).Error; err != nil {
		log.Error("error on updating user", err.Error())
		return domain.ContentCore{}, domain.User{}, err
	}

	return ToDomain(cnv), ToDomainU(res), nil
}

// GetAll implements domain.Repository
func (rq *repoQuery) GetAll() ([]domain.ContentCore, error) {
	var resQry []Content
	if err := rq.db.Table("contents").Select("contents.id", "contents.article", "contents.point_art", "contents.user_id", "users.fullname").Joins("join users on users.id = contents.user_id").Scan(&resQry).Find(&resQry).Error; err != nil {
		log.Error("Error on All user", err.Error())
		return nil, err
	}
	return ToDomainArray(resQry), nil
}

// GetMy implements domain.Repository
func (rq *repoQuery) GetMy(userID uint, key uint) ([]domain.ContentCore, error) {
	var resQry []Content
	if err := rq.db.Table("contents").Select("contents.id", "contents.article", "contents.point_art", "contents.user_id", "users.fullname").Joins("join users on users.id = contents.user_id").Scan(&resQry).Where("user_id = ? AND contents.id = ?", userID, key).Find(&resQry).Error; err != nil {
		log.Error("Error on All user", err.Error())
		return nil, err
	}
	return ToDomainArray(resQry), nil
}

func (rq *repoQuery) GetMyAll(userID uint) ([]domain.ContentCore, error) {
	var resQry []Content
	if err := rq.db.Table("contents").Select("contents.id", "contents.article", "contents.point_art", "contents.user_id", "users.fullname").Joins("join users on users.id = contents.user_id").Scan(&resQry).Where("user_id = ? ", userID).Find(&resQry).Error; err != nil {
		log.Error("Error on All user", err.Error())
		return nil, err
	}
	return ToDomainArray(resQry), nil
}
