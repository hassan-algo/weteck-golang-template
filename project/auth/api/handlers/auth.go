package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	extras "madeer.com/extras"
)

type AuthHanders struct {
	myData string
}

func NewAuthHandler(temp string) *AuthHanders {
	return &AuthHanders{myData: temp}
}

func (ar *AuthHanders) Get(c echo.Context) error {

	type res struct {
		Value string `json:"title"`
		age   int8   `json:"age"`
	}

	data := &res{
		Value: "Hello, Shoaib!",
		age:   21,
	}

	// return c.String(http.StatusOK, "Hello, World!")
	return c.JSON(http.StatusOK, data)
}

type res struct {
	Value string `json:"title"`
	Age   int    `json:"age"`
}

func (ar *AuthHanders) Post(c echo.Context) error {
	data, err := extras.GetJSONRawBody(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	var myData res
	if encodedData, err := json.Marshal(data); err == nil {
		if err := json.Unmarshal(encodedData, &myData); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusOK, myData)
	}

	// return c.String(http.StatusOK, "Hello, World!")
	return c.JSON(http.StatusBadRequest, nil)
}
