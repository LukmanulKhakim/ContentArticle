package services

import (
	"content/config"
	"content/feature/article/domain"
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type contentService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &contentService{
		qry: repo,
	}
}

// EditPoint implements domain.Service
func (cs *contentService) EditPoint(point domain.ContentCore, contentID uint, user uint) (domain.ContentCore, domain.User, error) {
	res, res2, err := cs.qry.Edit(point, contentID, user)
	if err != nil {
		return domain.ContentCore{}, domain.User{}, errors.New(config.DATABASE_ERROR)
	}
	return res, res2, nil
}

// GetAllContent implements domain.Service
func (cs *contentService) GetAllContent() ([]domain.ContentCore, error) {
	res, err := cs.qry.GetAll()
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

// GetMyContent implements domain.Service
func (cs *contentService) GetMyContent(userID uint, key uint) ([]domain.ContentCore, error) {
	res, err := cs.qry.GetMy(userID, key)
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

func (cs *contentService) GetMyAllContent(userID uint) ([]domain.ContentCore, error) {
	res, err := cs.qry.GetMyAll(userID)
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

// Post implements domain.Service
func (cs *contentService) Post(newItem domain.ContentCore) (domain.ContentCore, error) {
	newItem.Point_Art = 0
	res, err := cs.qry.Add(newItem)
	if err != nil {
		return domain.ContentCore{}, errors.New("some problem on database")
	}
	return res, nil
}
