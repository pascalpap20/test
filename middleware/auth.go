package middleware

import (
	"crud/dto"
	"crud/utils/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func IsSuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Perform operations before calling the next handler
		authHeader := c.GetHeader("Authorization")

		// Check if the header starts with "Bearer"
		if strings.HasPrefix(authHeader, "Bearer ") {
			// Extract the token value
			token := strings.TrimPrefix(authHeader, "Bearer ")

			// Perform operations with the token value
			claim, err := auth.VerifyToken(token)
			if err != nil {
				panic(err)
				return
			}

			if claim["role_id"] == float64(1) {
				// Call the next handler
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, dto.DefaultUnauthorizedResponse())
				c.Abort()
				return
			}
		}
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Perform operations before calling the next handler
		authHeader := c.GetHeader("Authorization")

		// Check if the header starts with "Bearer"
		if strings.HasPrefix(authHeader, "Bearer ") {
			// Extract the token value
			token := strings.TrimPrefix(authHeader, "Bearer ")

			// Perform operations with the token value
			claim, err := auth.VerifyToken(token)
			if err != nil {
				panic(err)
				return
			}

			if claim["role_id"] == float64(2) {
				// Call the next handler
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, dto.DefaultUnauthorizedResponse())
				c.Abort()
				return
			}
		}
	}
}

func IsSuperAdminOrAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Perform operations before calling the next handler
		authHeader := c.GetHeader("Authorization")

		// Check if the header starts with "Bearer"
		if strings.HasPrefix(authHeader, "Bearer ") {
			// Extract the token value
			token := strings.TrimPrefix(authHeader, "Bearer ")

			// Perform operations with the token value
			claim, err := auth.VerifyToken(token)
			if err != nil {
				panic(err)
				return
			}

			if claim["role_id"] == float64(1) || claim["role_id"] == float64(2) {
				// Call the next handler
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, dto.DefaultUnauthorizedResponse())
				c.Abort()
				return
			}
		}
	}
}
