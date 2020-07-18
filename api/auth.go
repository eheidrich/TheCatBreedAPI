package api

import (
	"net/http"
	"strings"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetToken(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	user := auth{}
	c.Bind(&user)
	ok, err := valid.Valid(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	// Validate a fixed user (could do a user model and service to validate, but it is out of scope)
	if user.Username != "admin" || user.Password != "@#$RF@!718" {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
	}

	token, err := util.GenerateToken(user.Username, user.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Authorization Token is required"})
			c.Abort()
			return
		}

		extractedToken := strings.Split(token, "Bearer ")[1]

		_, err := util.ParseToken(extractedToken)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Authorization Token timed out"})
			default:
				c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Authorization Token is invalid"})
			}

			c.Abort()
			return
		}

		c.Next()
	}
}
