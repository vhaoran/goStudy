package jwt

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/robbert229/jwt"
)

func Test_jwt_gen(t *testing.T) {
	secret := "ThisIsMySuperSecret"
	algorithm := jwt.HmacSha256(secret)

	claims := jwt.NewClaim()
	claims.Set("Role", "Admin")
	claims.Set("UserName", "whr")
	claims.Set("RoomID", "123")
	claims.SetTime("exp", time.Now().Add(time.Minute))

	token, err := algorithm.Encode(claims)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Token: %s\n", token)

	if algorithm.Validate(token) != nil {
		panic(err)
	}

	loadedClaims, err := algorithm.Decode(token)
	if err != nil {
		panic(err)
	}

	role, err := loadedClaims.Get("Role")
	if err != nil {
		panic(err)
	}

	userName, err := loadedClaims.Get("UserName")
	if err != nil {
		panic(err)
	}
	log.Println("userName:", userName)

	roomID, err := loadedClaims.Get("RoomID")
	if err != nil {
		panic(err)
	}
	log.Println("roomID:", roomID)

	roleString, ok := role.(string)
	if !ok {
		panic(err)
	}

	if strings.Compare(roleString, "Admin") == 0 {
		//user is an admin
		fmt.Println("User is an admin")
	}
}

func Test_do_jwt(t *testing.T) {
	secretKey := "abc"
	s, err := jwtGen(secretKey, int64(100))
	fmt.Println("------", "", "-----------")
	if err != nil {
		log.Println(err)
		return
	}

	//
	uid, err := jwtParse(secretKey, s)
	if err != nil {
		log.Println(err)
	}
	log.Println("uid", uid)
	fmt.Println("------", "ok", "-----------")
}

func jwtGen(secretKey string, uid int64) (string, error) {
	algorithm := jwt.HmacSha256(secretKey)

	claims := jwt.NewClaim()
	claims.Set("uid", fmt.Sprint(uid))
	//claims.SetTime("exp", time.Now().Add(time.Hour*24*365*10))

	token, err := algorithm.Encode(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func jwtParse(secretKey, token string) (uid int64, err error) {
	uid, err = 0, nil

	algorithm := jwt.HmacSha256(secretKey)
	//
	if err = algorithm.Validate(token); err != nil {
		return
	}

	//parse
	loadedClaims, err := algorithm.Decode(token)
	if err != nil {
		return 0, err
	}

	id, err := loadedClaims.Get("uid")
	if err != nil {
		return 0, err
	}

	s, ok := id.(string)
	if !ok {
		log.Println("uid:", uid)
		return 0, errors.New("错误的uid类型")
	}
	if uid, err = strconv.ParseInt(s, 10, 64); err != nil {
		log.Println("获取jwt时出错,err:", err)
		return 0, err
	}
	return
}
