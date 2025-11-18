package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从Header中获取Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权：缺少Authorization Header"})
			c.Abort()
			return
		}

		// 2. 检查Token格式
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权：Token格式错误"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 3. 解析Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 检查签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// 返回密钥
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			// 检查是否是过期错误
			if errors.Is(err, jwt.ErrTokenExpired) {
				c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权：Token已过期"})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": fmt.Sprintf("未授权：Token解析失败: %v", err)})
			c.Abort()
			return
		}

		// 4. 验证Token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 5. 将用户ID存入Context
			userID, ok := claims["user_id"].(float64)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权：Token中缺少用户ID"})
				c.Abort()
				return
			}
			c.Set("user_id", uint64(userID))
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权：Token无效"})
			c.Abort()
			return
		}
	}
}
