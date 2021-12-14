package pkg

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"sync"
	"time"
)

type Jwt struct {
}

func (j *Jwt) createToken(uid int64, secret string) (string, error) {
	claim := jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(3 * 24 * time.Hour).Unix(),
	}
	nt := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return nt.SignedString([]byte(secret))
}

func (j *Jwt) refreshToken(tokenStr string, secret string) (string, error) {
	claim, err := j.parseToken(tokenStr, secret)
	if claim == nil || err != nil {
		err = errors.New("parse token error")
		return "", err
	}
	uid := claim["uid"]
	return j.createToken(uid.(int64), secret)
}

func (j *Jwt) parseToken(tokenStr string, secret string) (map[string]interface{}, error) {
	getSecret := func(_secret string) func(token *jwt.Token) (i interface{}, err error) {
		return func(token *jwt.Token) (i interface{}, err error) {
			return []byte(_secret), nil
		}
	}(secret)
	token, err := jwt.Parse(tokenStr, getSecret)
	if err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapClaim")
		return nil, err
	}
	if !token.Valid {
		err = errors.New("token is invalid")
		return nil, err
	}
	return claim, nil
}

var (
	onceJwt     sync.Once
	instanceJwt *Jwt
)

func GetJwtInstance() *Jwt {
	onceJwt.Do(func() {
		instanceJwt = &Jwt{}
	})
	return instanceJwt
}

func GenJwtToken(uid int64, secret string) string {
	j := GetJwtInstance()
	token, err := j.createToken(uid, secret)
	if err != nil {
		panic(err)
	}
	return token
}

func VerifyJwtToken(tokenStr string, secret string) bool {
	j := GetJwtInstance()
	_, err := j.parseToken(tokenStr, secret)
	if err != nil {
		return false
	}
	return true
}

func ParseJwtToken(tokenStr string, secret string) (map[string]interface{}, error) {
	j := GetJwtInstance()
	claim, err := j.parseToken(tokenStr, secret)
	return claim, err
}
