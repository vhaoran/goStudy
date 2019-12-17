package jwt

import (
	"fmt"
	"log"
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
