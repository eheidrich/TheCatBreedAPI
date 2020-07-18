package main

import (
	"github.com/eheidrich/TheCatBreedAPI/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.POST("/auth", api.GetToken)
	r.GET("/breeds", api.ValidateToken(), api.SearchCatBreads)

	r.Run()
}
