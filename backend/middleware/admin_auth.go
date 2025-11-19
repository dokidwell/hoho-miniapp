package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AdminAuthMiddleware 管理员JWT认证中间件
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权：缺少Authorization Header"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权：Token格式错误"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权：Token已过期"})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": fmt.Sprintf("未授权：Token解析失败: %v", err)})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			adminID, ok := claims["admin_id"].(float64)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权：Token中缺少管理员ID"})
				c.Abort()
				return
			}
			c.Set("admin_id", uint64(adminID))
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权：Token无效"})
			c.Abort()
			return
		}
	}
}
