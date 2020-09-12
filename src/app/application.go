package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pgrau/bookstore-oauth-api/src/service/access_token"
	"github.com/pgrau/bookstore-oauth-api/src/http"
	"github.com/pgrau/bookstore-oauth-api/src/repository/db"
	"github.com/pgrau/bookstore-oauth-api/src/repository/rest"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(rest.NewRestUserRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}