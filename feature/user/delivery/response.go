package delivery

import "content/feature/user/domain"

type LoginResponse struct {
	Fullname string `json:"fullname"`
	Role     uint   `json:"role"`
	Token    string `json:"token"`
}

type RegistResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Point    int    `json:"point"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "login":
		cnv := core.(domain.UserCore)
		res = LoginResponse{Fullname: cnv.Fullname, Role: cnv.Role, Token: cnv.Token}
	case "reg":
		cnv := core.(domain.UserCore)
		res = RegistResponse{Fullname: cnv.Fullname, Email: cnv.Email}
	case "user":
		cnv := core.(domain.UserCore)
		res = UserResponse{ID: cnv.ID, Fullname: cnv.Fullname, Email: cnv.Email, Point: cnv.Point}
	case "all":
		var arr []UserResponse
		cnv := core.([]domain.UserCore)
		for _, val := range cnv {
			arr = append(arr, UserResponse{ID: val.ID, Fullname: val.Fullname, Email: val.Email, Point: val.Point})
		}
		res = arr
	}
	return res
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func SuccessLogin(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func SuccessDeleteResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}
