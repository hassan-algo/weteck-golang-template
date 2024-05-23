package business

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"

	"example.com/db"
	"example.com/extras"
	"example.com/structs"
)

type AuthBusiness struct {
	dbCon *db.DatabaseConnection
}

func NewAuthBusiness() *AuthBusiness {
	return &AuthBusiness{}
}
func (b *AuthBusiness) Connect(dbConnection *db.DatabaseConnection) error {
	b.dbCon = dbConnection
	return nil
}
func (b *AuthBusiness) Authentication(email string, password string) (interface{}, error) {

	hash := md5.Sum([]byte(password))
	hashPassword := hex.EncodeToString(hash[:])

	query := fmt.Sprintf("SELECT userguid, fullname, profilepic,email,password, usertype, login_token  FROM tbl_users where (LOWER(email) = LOWER('%s') or LOWER(contact) = LOWER('%s')) AND password = '%s'", email, email, hashPassword)
	rowsRs, err := b.dbCon.Con.Query(query)

	resErr := structs.Response{
		Valid:   false,
		Message: "Auth Failed!",
		Data:    nil,
	}

	if err != nil {
		resErr.Message = "Auth Failed!" + err.Error()
		return resErr, err
	}
	defer rowsRs.Close()

	// creates placeholder of the Credentials
	results := make([]structs.Credentials, 0)

	// we loop through the values of rows
	for rowsRs.Next() {
		obj := structs.Credentials{}
		// err := rowsRs.Scan(&snb.FullName, &snb.OwnerGuid, &snb.Contact, &snb.PIN, &snb.ProfilePic, &snb.Email, &snb.Password)
		err := rowsRs.Scan(&obj.UserGuid, &obj.FullName, &obj.ProfilePic, &obj.Email, &obj.Password, &obj.UserType, &obj.Login_Token)
		if err != nil {
			resErr.Message = err.Error()
			return resErr, err

		}
		results = append(results, obj)
	}

	if err = rowsRs.Err(); err != nil {
		resErr.Message = err.Error()
		return resErr, err

	}

	// result is array of objects
	if len(results) < 1 {
		resErr.Message = "Data not received"
		return resErr, err

	} else {

		secretKey := extras.GetSecretKey()

		secretKeyQuery := `UPDATE tbl_users SET login_token = $1 WHERE userguid = $2;`
		_, err := b.dbCon.Con.Exec(secretKeyQuery, secretKey, results[0].UserGuid)

		if err != nil {
			resErr.Message = err.Error()
			// extras.LogThisWithActor(i.e, "Update Secret Key: "+err.Error(), "")
		}

		// make jwt token
		claims := jwt.MapClaims{}
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token will expire in 24 hours
		claims["userguid"] = results[0].UserGuid
		claims["email"] = results[0].Email

		// Create the token using the claims and a secret key
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := token.SignedString([]byte(secretKey))

		if err != nil {
			resErr.Message = err.Error()
			return resErr, err

		}

		resUser := structs.ResponseUserWithToken{
			Name:       results[0].FullName,
			ProfilePic: results[0].ProfilePic,
			Email:      results[0].Email,
			Token:      signedToken + " " + results[0].UserGuid,
		}

		res := structs.Response{
			Valid:   true,
			Message: results[0].UserType,
			Data:    resUser,
		}

		return res, nil
	}
}

func (b *AuthBusiness) Authenticate(userGuid string, token string) (error, string, string) {

	var (
		JWT_KEY         string
		updatedUserGuid string
		role            string
	)

	//
	err := b.dbCon.Con.QueryRow("SELECT login_token, userguid, usertype FROM tbl_users WHERE userguid = $1", userGuid).Scan(&JWT_KEY, &updatedUserGuid, &role)
	if err == sql.ErrNoRows {
		// extras.LogThisWithActor(i.e, "Can't get any rows", "") //
		return errors.New("auth Failed2"), "", ""
	} else if err != nil {
		// extras.LogThisWithActor(i.e, err.Error(), "")
		return errors.New("server Error"), "", ""
	}

	// extras.LogThisWithActor(i.e, "jwt:"+JWT_KEY, "")           //
	// extras.LogThisWithActor(i.e, "uguid:"+updatedUserGuid, "") //
	// extras.LogThisWithActor(i.e, "role:"+role, "")             //

	if JWT_KEY != "" {
		claims := &jwt.StandardClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_KEY), nil
		})
		if err != nil {
			// extras.LogThisWithActor(i.e, err.Error(), "") //

			return errors.New("auth Failed3"), "", ""
		}

		// extras.LogThisWithActor(i.e, "", "Candidate")
		return nil, updatedUserGuid, role //
	} else {
		// extras.LogThisWithActor(i.e, err.Error(), "")
		return errors.New("auth Failed5"), "", ""
	}
}
