package handler

import (
	"net/http"

	"github.com/SunilKividor/PillNet-Backend/internal/authentication/auth"
	"github.com/gin-gonic/gin"
)

type AuthenticationHandler struct {
	Authentication auth.Authentication
}

func NewAuthenticationHandler(auth auth.Authentication) *AuthenticationHandler {
	return &AuthenticationHandler{
		Authentication: auth,
	}
}

type AuthLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthSignUpReq struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AuthRefreshReq struct {
	RefreshToken string `json:"refresh_token"`
}

func (ah *AuthenticationHandler) Login(c *gin.Context) {
	var reqB AuthLoginReq

	if c.ShouldBindJSON(&reqB) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request format"})
		return
	}

	ctx := c.Request.Context()

	authTokens, err := ah.Authentication.Login(ctx, reqB.Username, reqB.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authTokens)
}

func (ah *AuthenticationHandler) SignUp(c *gin.Context) {
	var reqB AuthSignUpReq

	if c.ShouldBindJSON(&reqB) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request format"})
		return
	}

	ctx := c.Request.Context()

	authTokens, err := ah.Authentication.SignUp(ctx, reqB.Name, reqB.Username, reqB.Password, reqB.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authTokens)
}

func (ah *AuthenticationHandler) Refresh(c *gin.Context) {
	var reqB AuthRefreshReq

	if c.ShouldBindJSON(&reqB) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request format"})
		return
	}

	ctx := c.Request.Context()

	authTokens, err := ah.Authentication.Refresh(ctx, reqB.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authTokens)
}

func (ah *AuthenticationHandler) Logout(c *gin.Context) {

	id := c.GetString("id")

	ctx := c.Request.Context()

	err := ah.Authentication.Logout(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
