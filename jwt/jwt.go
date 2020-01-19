/**
 * Created by Goland.
 * User: TinyStar
 * Date: 2019/11/11
 * Time: 10:56
 */
package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_"log"
	"time"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserID      int64     `json:"user_id"`
}

var Pl JWTClaims

func MakeToke(id int64)  string{

	claims := JWTClaims{
		UserID:id,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt=time.Now().Add(time.Hour * time.Duration(1)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("sign.verify"))
	if err != nil {
		//log.Fatalln(err)
	}
	return signedToken
}

func ValidateToken(token1 string) bool{
	authorization := token1
	token,err := jwt.Parse(authorization, func(token *jwt.Token) (i interface{}, e error) {
		return []byte("sign.verify"),nil
	})

	if err != nil {
		if err ,ok := err.(*jwt.ValidationError);ok {
			if err.Errors & jwt.ValidationErrorMalformed != 0 {
				return false
			}
			if err.Errors & (jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
				fmt.Println(err)
				return false
			}
		}
		return false
	}

	if err :=token.Claims.Valid();err==nil{
		finToken := token.Claims.(jwt.MapClaims)

		Pl.UserID = int64(finToken["user_id"].(float64))
		return true
	}
	return false
}