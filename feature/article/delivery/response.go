package delivery

import "content/feature/article/domain"

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func SuccessDeleteResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func FailResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

type AddResponse struct {
	ID        uint   `json:"id"`
	Article   string `json:"article"`
	Point_Art int    `json:"point_art"`
	User_ID   uint   `json:"user_id"`
}

type GetResponse struct {
	ID        uint   `json:"id"`
	Article   string `json:"article"`
	Point_Art int    `json:"point_art"`
	User_ID   uint   `json:"user_id"`
	Fullname  string `json:"fullname"`
}

type UpdateResponse struct {
	ID        uint   `json:"id_product"`
	Point_Art int    `json:"point_art"`
	User_ID   uint   `json:"user_id"`
	Fullname  string `json:"fullname"`
}

func ToResponse(basic interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "add":
		cnv := basic.(domain.ContentCore)
		res = AddResponse{ID: cnv.ID, Article: cnv.Article, Point_Art: cnv.Point_Art, User_ID: cnv.User_ID}
	case "point":
		cnv := basic.(domain.ContentCore)
		res = UpdateResponse{ID: cnv.ID, Point_Art: cnv.Point_Art, User_ID: cnv.User_ID, Fullname: cnv.Fullname}
	case "all":
		var arr []GetResponse
		cnv := basic.([]domain.ContentCore)
		for _, val := range cnv {
			arr = append(arr, GetResponse{ID: val.ID, Article: val.Article, Point_Art: val.Point_Art, User_ID: val.User_ID, Fullname: val.Fullname})
		}
		res = arr
	}
	return res
}
