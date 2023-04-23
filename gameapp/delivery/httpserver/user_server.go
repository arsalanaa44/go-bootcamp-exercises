package httpserver

import (
	"gameapp/service/userservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) userRegister(c echo.Context) error {

	var uReq userservice.RegisterRequest
	if bErr := c.Bind(&uReq); bErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, bErr.Error())
	}

	if uRes, lErr := s.userSvc.Register(uReq); lErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, lErr.Error())
	} else {

		return c.JSON(http.StatusCreated, uRes)
	}

}
