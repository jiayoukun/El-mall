package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("JWT_SECRET")

type Claims struct {
	Id uint `json:"id"`
	Username string	`json:"username"`
	Password string `json:"password"`
	Authority int `json:"authority"`
	jwt.StandardClaims
}

//签发用户token
func Generatetoken(id uint,username,password string,authority int)(string,error)  {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		id,
		username,
		password,
		authority,
		jwt.StandardClaims{
			ExpiresAt:expireTime.Unix(),
			Issuer:"el_mall",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	token,err:=tokenClaims.SignedString(jwtSecret)
	return token,err
}

//校验用户token
func Checktoken(token string)(*Claims,error)  {
	tokenclaims,err:=jwt.ParseWithClaims(token,&Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret,nil
	})
	if tokenclaims != nil {
		if claims,ok := tokenclaims.Claims.(*Claims);ok && tokenclaims.Valid {
			return claims,nil
		}
	}
	return nil, err
}