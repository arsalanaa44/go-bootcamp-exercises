package httpserver

import (
	"fmt"
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
func (s Server) userLogin(c echo.Context) error {

	var lReq userservice.LoginRequest
	if bErr := c.Bind(&lReq); bErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, bErr.Error())
	}

	if lRes, lErr := s.userSvc.Login(lReq); lErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, lErr.Error())
	} else {

		return c.JSON(http.StatusOK, lRes)
	}

}

func (s Server) userProfile(c echo.Context) error {

	tokenString := c.Request().Header.Get("Authorization")
	fmt.Println(tokenString)
	claims, pErr := s.authSvc.ParseToken(tokenString)

	fmt.Println(s.authSvc.ParseToken(tokenString))
	if pErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, pErr.Error())
	}
	req := userservice.ProfileRequest{claims.UserID}

	if res, err := s.userSvc.Profile(req); err != nil {

		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	} else {

		return c.JSON(http.StatusOK, res)
	}

}

//// reads json file well
//func (s Server) userProfile(c echo.Context) error {
//	var pReq userservice.ProfileRequest
//	if err := c.Bind(&pReq); err != nil {
//
//		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//	}
//
//	return s.userHandler(c, pReq, func(req interface{}) (interface{}, error) {
//		pReq, ok := req.(userservice.ProfileRequest)
//		if !ok {
//
//			return nil, fmt.Errorf("invalid request type")
//		}
//		return s.userSvc.Profile(pReq)
//	})
//}
//
//type any = interface{}
//
//func (s Server) userHandler(c echo.Context, req any, f func(any) (any, error)) error {
//	if res, lErr := f(req); lErr != nil {
//
//		return echo.NewHTTPError(http.StatusBadRequest, lErr.Error())
//	} else {
//
//		return c.JSON(http.StatusOK, res)
//	}
//}
