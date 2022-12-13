package delivery

import (
	ck "content/config"
	"content/feature/article/domain"
	"content/utils/common"
	"errors"
	lo "log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type contentHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := contentHandler{srv: srv}
	e.POST("/content", handler.Post(), middleware.JWT([]byte(ck.JwtKey)))
	e.GET("/admin/content", handler.GetAllContent(), middleware.JWT([]byte(ck.JwtKey)))
	e.GET("/content", handler.GetMyAllContent(), middleware.JWT([]byte(ck.JwtKey)))
	e.GET("/content/:id", handler.GetMyContent(), middleware.JWT([]byte(ck.JwtKey)))
	e.PUT("/point/:id", handler.UpdatePoint(), middleware.JWT([]byte(ck.JwtKey)))
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

func (ch *contentHandler) GetAllContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, role := common.ExtractToken(c)
		if role != 1 {
			return c.JSON(http.StatusUnauthorized, FailResponse("this account not admin"))
		} else if userID == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		}
		res, err := ch.srv.GetAllContent()
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				c.JSON(http.StatusBadRequest, FailResponse(err.Error()))
			} else {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		} else {
			return c.JSON(http.StatusOK, SuccessResponse("success get all content by admin", ToResponse(res, "all")))
		}
		return nil
	}
}

func (ch *contentHandler) GetMyContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, role := common.ExtractToken(c)
		if role != 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("this account for user"))
		} else if userID == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return errors.New("cannot convert id")
		}
		res, err := ch.srv.GetMyContent(userID, uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("An invalid client request"))
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success show my content", ToResponse(res, "all")))
	}
}

func (ch *contentHandler) GetMyAllContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, role := common.ExtractToken(c)
		if role != 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("this account for user"))
		} else if userID == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		}
		res, err := ch.srv.GetMyAllContent(userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("An invalid client request"))
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success show my content", ToResponse(res, "all")))
	}
}

func (ch *contentHandler) UpdatePoint() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		userID, role := common.ExtractToken(c)
		if role != 1 {
			return c.JSON(http.StatusUnauthorized, FailResponse("this account for user"))
		} else if userID == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		} else {
			ID, err := strconv.Atoi(c.Param("id"))
			if err := c.Bind(&input); err != nil {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
			}
			if err != nil {
				return c.JSON(http.StatusBadRequest, FailResponse("id poli must integer"))
			}
			cnv := ToDomain(input)
			res, res2, err := ch.srv.EditPoint(cnv, uint(ID), cnv.User_ID)
			lo.Println(res2)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
			return c.JSON(http.StatusCreated, SuccessResponse("sucses edit poli", ToResponse(res, "edit")))

		}

	}
}
