package crypto

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/z13z/Kiosks/master-server/db/users"
	"log"
	"strconv"
	"time"
)

//todo zaza
var signatureKey = []byte("changeIt")
var durationForKey, _ = time.ParseDuration("1h")

type KioskUserClaims struct {
	Username    string `json:"username"`
	Permissions string `json:"permissions"`
	jwt.StandardClaims
}

func GetJwtForUser(entity users.UserEntity) (string, error) {
	val, err := entity.Permissions.Value()
	if err != nil {
		return "", err
	}
	permissions := val.(string)
	userClaims := KioskUserClaims{
		entity.Name,
		permissions,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(durationForKey).Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims).SignedString(signatureKey)
}

func CheckJwtForUser(tokenToCheck string) (*KioskUserClaims, bool) {
	if tokenToCheck == "" {
		return nil, false
	}
	token, err := jwt.ParseWithClaims(tokenToCheck, &KioskUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signatureKey, nil
	})
	if err != nil {
		log.Print("Can't parse jwt: "+tokenToCheck+": ", err)
		return nil, false
	}
	parsedClaims, ok := token.Claims.(*KioskUserClaims)
	if token.Valid && ok {
		return parsedClaims, ok
	} else {
		return nil, false
	}
}

func GetJwtForKiosk(id int64) (string, error) {
	claims := jwt.StandardClaims{
		Id:       strconv.FormatInt(id, 10),
		IssuedAt: time.Now().Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(signatureKey)
}

func CheckJwtForKiosk(tokenToCheck string) bool {
	if tokenToCheck == "" {
		return false
	}
	token, err := jwt.ParseWithClaims(tokenToCheck, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signatureKey, nil
	})
	if err != nil {
		log.Print("Can't parse jwt: "+tokenToCheck, err)
		return false
	}
	_, ok := token.Claims.(*jwt.StandardClaims)
	if token.Valid && ok {
		return ok
	} else {
		return false
	}
}
