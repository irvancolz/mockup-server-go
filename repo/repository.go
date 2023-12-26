package repo

import (
	"log"
	helper "mockup_server/helper"
	"mockup_server/model"

	"gopkg.in/dgrijalva/jwt-go.v3"
)

func SendLoginReq(username, pass string) (*model.User, error) {
	var result model.User

	dbConn := helper.InitDbConn()
	query := `
	UPDATE public.user set 
		is_login = true
	WHERE username = $1
	AND password = $2
	AND is_login = false
	`

	_, errExec := dbConn.Exec(query, username, pass)
	if errExec != nil {
		log.Println("failed to login :", errExec)
		return nil, errExec
	}

	getUserDataQuery := `
	SELECT 
		username,
		is_login,
		role
	FROM public.user WHERE username = $1`

	rows := dbConn.QueryRowx(getUserDataQuery, username)

	if errScan := rows.StructScan(&result); errScan != nil {
		log.Println("failed read user data :", errScan)
		return nil, errScan
	}

	return &result, nil
}

func SendLogoutReq(id string) (int64, error) {
	dbConn := helper.InitDbConn()
	query := `
	UPDATE public.user set 
		is_login = false
	WHERE id = $1
	AND is_login = true
	`

	execRes, errExec := dbConn.Exec(query, id)
	if errExec != nil {
		log.Println("failed to logout :", errExec)
		return 0, errExec
	}

	result, errGetAffectedRes := execRes.RowsAffected()
	if errGetAffectedRes != nil {
		log.Println("failed to loged out user :", errGetAffectedRes)
	}

	return result, nil
}

func GenerateJWT(data model.User) string {

	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	signedToken, erreSign := tokens.SignedString([]byte("lorem"))
	if erreSign != nil {
		log.Println("failed sign token :", erreSign)
		return ""
	}

	return signedToken
}
