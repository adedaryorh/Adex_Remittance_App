package api

import "C"
import (
	db "Fin-Remittance/db/sqlc"
	"Fin-Remittance/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"time"
)

type User struct {
	server *Server
}

func (u User) router(server *Server) {
	u.server = server

	serverGroup := server.router.Group("/users")
	serverGroup.GET("", u.listUsers)
	serverGroup.GET("", u.listUsers)
	serverGroup.POST("me", u.getLoggedInUser)
}

type UserParams struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func (u *User) createUser(c *gin.Context) {
	//creating a user instance
	user := new(UserParams)
	//or var user UserParams
	// binding json payload
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := utils.GenerateHashedPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateUserParams{
		Email:          user.Email,
		HashedPassword: hashedPassword,
		Username:       user.Username,
	}

	newUser, err := u.server.queries.CreateUser(context.Background(), arg)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
				return
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, UserResponse{}.toUserResponse(&newUser))
}

func (u *User) listUsers(c *gin.Context) {
	arg := db.ListUserParams{
		Offset: 0,
		Limit:  10,
	}
	users, err := u.server.queries.ListUser(context.Background(), arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (u User) getLoggedInUser(c *gin.Context) {
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to access resources"})
	}

	converted, ok := userId.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Encountered an issue"})
	}
}

type UserResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u UserResponse) toUserResponse(user *db.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

}
