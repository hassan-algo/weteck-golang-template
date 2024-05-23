package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"example.com/apis"
	"example.com/extras"
	"example.com/structs"
)

type AuthHandlers struct {
	authBusiness apis.AuthBusiness
}

func NewAuthHandler() *AuthHandlers {
	return &AuthHandlers{}
}

func (h *AuthHandlers) Authentication(ec echo.Context) error {
	body := extras.GetJSONRawBody(ec)
	// extras.LogThisWithActor(i.e, "", body["email"].(string))

	email := body["email"].(string)
	password := body["password"].(string)

	data, err := h.authBusiness.Authentication(email, password)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, data)
	}
	return ec.JSON(http.StatusOK, data)
}

// Authenticate wraps a function with authentication logic.
func (h *AuthHandlers) Authenticate(f func(ec echo.Context) error, role ...string) func(ec echo.Context) error {
	return func(ec echo.Context) error {
		// get headers
		authHeader := ec.Request().Header.Get("Authorization")
		if authHeader == "" {
			return errors.New("auth Failed1")
		}
		splitHeader := strings.Split(authHeader, " ")

		userGuid := splitHeader[2]
		token := splitHeader[1]

		err, userGuid, returnedRole := h.authBusiness.Authenticate(userGuid, token)

		if err != nil {

			res := structs.Response{
				Valid:   false,
				Message: "UnAuthorized Request1",
				Data:    nil,
			}
			return ec.JSON(http.StatusUnauthorized, res)
		}

		if !extras.Contains(role, returnedRole) {
			res := structs.Response{
				Valid:   false,
				Message: "UnAuthorized Request2",
				Data:    nil,
			}

			return ec.JSON(http.StatusUnauthorized, res)
		}

		ec.Set("user_guid", userGuid)
		// Proceed with the original function if authentication is successful
		return f(ec)
	}
}

// func (h *AuthHandlers) CheckAuth(ec echo.Context) error {
// 	authHeader := ec.Request().Header.Get("Authorization")
// 	if authHeader == "" {
// 		return errors.New("Authentication failed, missing header")
// 	}

// 	splitHeader := strings.Split(authHeader, " ")
// 	if len(splitHeader) != 2 || splitHeader[0] != "Bearer" {
// 		return errors.New("Authentication failed, invalid header format")
// 	}

// 	token := splitHeader[1]
// 	claims := &jwt.StandardClaims{}

// 	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(h.authBusiness.GetSecretKey()), nil
// 	})
// 	if err != nil {
// 		return errors.New("Authentication failed: " + err.Error())
// 	}

// 	userData, err := h.authBusiness.GetUserDetails(claims.Subject)
// 	if err != nil {
// 		return errors.New("Authentication failed: " + err.Error())
// 	}

//		ec.Set("userData", userData)
//		return nil
//	}
//
// Connect is required to fulfill the apis.AuthHandler interface
func (h *AuthHandlers) Connect(business apis.AuthBusiness) error {
	h.authBusiness = business
	return nil
}

func (h *AuthHandlers) Middleware(f func(ec echo.Context) error) func(ec echo.Context) error {
	return func(ec echo.Context) error {
		// if err := h.CheckAuth(ec); err != nil {
		// 	res := structs.Response{
		// 		Valid:   false,
		// 		Message: err.Error(),
		// 		Data:    nil,
		// 	}
		// 	return ec.JSON(http.StatusForbidden, res)
		// }
		return f(ec)
	}
}

func (h *AuthHandlers) Decorate(f func(ec echo.Context) error) func(ec echo.Context) error {
	return func(ec echo.Context) error {
		return f(ec)
	}
}
