package extras

import (
	"encoding/json"
	"errors"

	"github.com/labstack/echo"
)

func GetJSONRawBody(c echo.Context) (map[string]interface{}, error) {

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		newErr := errors.New("Unable to decode JSON Request Body\n" + err.Error())
		return nil, newErr
	}

	return jsonBody, nil
}
