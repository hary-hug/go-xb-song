package util

import (
	"github.com/dgrijalva/jwt-go"
	"go-xb-song/applications/app/pkg/conf"
	"log"
	"time"
)

// define a custom claims
type Claims struct {
	Uid int
	jwt.StandardClaims
}

type Token struct {
	AccessToken  string  `json:"access_token"`
	ExpireAt     int64   `json:"expire_at"`
}

var jwtSecret []byte
var jwtTimeout int


func init()  {

	// get config
	sec, err := conf.IniFile.GetSection("app")
	if err != nil {
		log.Fatalln("Fail to get section 'app': ", err)
	}

	secret := sec.Key("JWT_SECRET").MustString("")
	timeout := sec.Key("JWT_TIMEOUT").MustInt(60)

	jwtSecret = []byte(secret)
	jwtTimeout = timeout
}

// creating a token using a custom claims type
func GenerateToken(uid int) (res Token, err error) {

	// expire time for token
	expire := time.Now().Add(time.Duration(jwtTimeout) * time.Second)

	claims := Claims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)

	res.AccessToken = tokenString
	res.ExpireAt = expire.Unix()

	return

}

// parse token using a custom claims
// return claims if success
func ParseToken(token string) (*Claims, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}