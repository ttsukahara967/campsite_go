package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("your_secret_key") // 本番は.envから読み込むなど

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

// ErrorResponse はエラーレスポンス用
type ErrorResponse struct {
	Error string `json:"error"`
}

// 仮のユーザー認証
func authenticateUser(username, password string) bool {
	return username == "admin" && password == "password"
}

// LoginHandler godoc
// @Summary ログイン
// @Description ユーザー名とパスワードでログイン
// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "ログイン情報"
// @Success 200 {object} LoginResponse
// @Failure 401 {object} ErrorResponse
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}
	// 認証チェック
	if req.Username != "admin" || req.Password != "password" {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "unauthorized"})
		return
	}

	// ★ JWT生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		// 必要なら "exp" など追加
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "token error"})
		return
	}
	c.JSON(http.StatusOK, LoginResponse{Token: tokenString})
}
