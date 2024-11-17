package app

import (
	accesstoken "BOOKSTORE_OAUTH-API/src/domain/access_token"
	"BOOKSTORE_OAUTH-API/src/http"
	"BOOKSTORE_OAUTH-API/src/repository/db"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(accesstoken.NewService(db.New()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")

}
