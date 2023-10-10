package api

import (
	db "Fin-Remittance/db/sqlc"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	queries *db.Queries
	router  *gin.Engine
}

func NewServer(port int) {
	g := gin.Default()

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "welcome to the Application"})
	})

	//g.Run(fmt.Sprintf(":"))
	g.Run(fmt.Sprintf(":%v", port))

}
