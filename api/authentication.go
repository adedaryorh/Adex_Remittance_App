package api

import (
	"Fin-Remittance/utils"
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Authentication struct {
	server *Server
}

func (a Authentication) router(server *Server) {
	a.server = server

	serverGroup := server.router.Group("/authentication")
	serverGroup.POST("login", a.login)
}

func (a Authentication) login(c *gin.Context) {
	user := new(UserParams)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := a.server.queries.GetUserByEmail(context.Background(), user.Email)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email or pass"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := utils.VerifyPassword(user.Password, dbUser.HashedPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email or pass"})
		return
	}
	token, err := utils.CreateToken(dbUser.ID, a.server.config.Signing_key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
