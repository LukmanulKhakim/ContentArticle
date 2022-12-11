package delivery

import (
	ck "content/config"
	"content/feature/article/domain"
	"content/utils/common"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type contentHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := contentHandler{srv: srv}
	e.POST("/content", handler.Post(), middleware.JWT([]byte(ck.JwtKey)))

}

func (ch *contentHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddFormat
		userID, role := common.ExtractToken(c)
		if role != 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("this account not user"))
		} else if userID == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		}
		input.User_ID = userID
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := ch.srv.Post(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("There is problem on server."))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("success add article", ToResponse(res, "add")))
	}

}
