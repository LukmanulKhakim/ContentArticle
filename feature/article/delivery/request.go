package delivery

import "content/feature/article/domain"

type AddFormat struct {
	Article string `json:"article" form:"article"`
	User_ID uint   `json:"user_id" form:"user_id"`
}

type UpdateFormat struct {
	ID        uint `json:"id" form:"id"`
	Point_Art int  `json:"point_art" form:"point_art"`
	User_ID   uint `json:"user_id" form:"user_id"`
}

func ToDomain(i interface{}) domain.ContentCore {
	switch i.(type) {
	case AddFormat:
		cnv := i.(AddFormat)
		return domain.ContentCore{Article: cnv.Article, User_ID: cnv.User_ID}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.ContentCore{ID: cnv.ID, Point_Art: cnv.Point_Art, User_ID: cnv.User_ID}
	}
	return domain.ContentCore{}
}
