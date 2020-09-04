package Controllers

import (
	"example.com/m/v2/Models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
)

func Login(c *gin.Context) {
	var user, userExist Models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	// compare the user form the rq, with the one we defined

	err := Models.FindByUserName(&userExist, user.UserName)
	if err != nil || user.PassWord != userExist.PassWord {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := createToken(userExist.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

func TokenValid(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func GetUserIdFromRequest(r *http.Request) string {
	token, _ := verifyToken(r)
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if userId, ok := claims["user_id"]; ok {
			return fmt.Sprint(userId)
		}
	}
	return ""
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	os.Setenv("ACCESS_SECRET", "SNMHzwzQIfdVYu3d4ZW03CpoMRuNh1rq")
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil

	})
	if err != nil {
		return nil, err
	}
	return token, nil

}

func createToken(userId uint) (string, error) {
	var err error

	// Create Access Token
	os.Setenv("ACCESS_SECRET", "SNMHzwzQIfdVYu3d4ZW03CpoMRuNh1rq")

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return "Bearer " + token, nil
}

func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
