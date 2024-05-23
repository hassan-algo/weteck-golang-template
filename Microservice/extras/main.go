package extras

import (
	"encoding/json"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func GetEnv(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ConvertDashesToUnderscores(input string) string {
	return strings.ReplaceAll(input, "-", "_")
}

func GetTypeForColumn(datatype string) string {
	if datatype == "Short Text" {
		return "VARCHAR(100)"
	} else if datatype == "Paragraph" {
		return "TEXT"
	} else if datatype == "Multiple choice" {
		return "TEXT"
	} else if datatype == "Yes/No" {
		return "INT"
	} else if datatype == "Checkbox" {
		return "TEXT"
	} else if datatype == "File upload" {
		return "TEXT"
	} else if datatype == "Multiple choice grid" {
		return "TEXT"
	} else if datatype == "Date" {
		return "DATE"
	} else if datatype == "Time" {
		return "TIME"
	} else if datatype == "Phone number" {
		return "VARCHAR(100)"
	} else if datatype == "Address" {
		return "VARCHAR(500)"
	} else if datatype == "Location" {
		return "VARCHAR(500)"
	} else if datatype == "Document" {
		return "TEXT"
	} else if datatype == "End screen" {
		return "INT"
	}
	return ""
}

// get secretkey
func GetSecretKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 8
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GetJSONRawBody(c echo.Context) map[string]interface{} {

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {

		// log.Error("empty json body")
		return nil
	}

	return jsonBody
}
