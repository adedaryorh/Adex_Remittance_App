package api

import (
	db "Fin-Remittance/db/sqlc"
	"Fin-Remittance/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

type Server struct {
	queries *db.Queries
	router  *gin.Engine
	config  *utils.Config
}

var tokenController *utils.JWTToken

func NewServer(envPath string) *Server {
	config, err := utils.LoadConfig(envPath)
	if err != nil {
		panic(fmt.Sprintf("Can not load env config: %v", err))
	}
	connection, err := sql.Open(config.DBdriver, config.DB_source_live)
	if err != nil {
		panic(fmt.Sprintf("could not load config: %v", err))
	}

	tokenController = utils.NewJWTToken(config)

	q := db.New(connection)
	g := gin.Default()

	return &Server{
		queries: q,
		router:  g,
		config:  config,
	}
	/*
		g := gin.Default()

		g.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "welcome to the Application"})
		})

		//g.Run(fmt.Sprintf(":"))
		g.Run(fmt.Sprintf(":%v", port))
	*/

}
func (s *Server) Start(port int) {
	s.router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "welcome to the Application"})
	})

	User{}.router(s)
	Authentication{}.router(s)
	s.router.Run(fmt.Sprintf(":%v", port))
}
