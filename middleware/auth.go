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

			if claim["role_id"] == float64(1) && claim["is_verified"] == "true" && claim["is_active"] == "true" {
				// Call the next handler
				c.Set("id", claim["id"])
				c.Set("role_id", claim["role_id"])
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, dto.DefaultUnauthorizedResponse())
				c.Abort()
				return
			}
		}
	}
}

//func IsAdmin() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// Perform operations before calling the next handler
//		authHeader := c.GetHeader("Authorization")
//
//		// Check if the header starts with "Bearer"
//		if strings.HasPrefix(authHeader, "Bearer ") {
//			// Extract the token value
//			token := strings.TrimPrefix(authHeader, "Bearer ")
//
//			// Perform operations with the token value
//			claim, err := auth.VerifyToken(token)
//			if err != nil {
//				panic(err)
//				return
//			}
//
//			if claim["role_id"] == float64(2) && claim["is_verified"] == "true" && claim["is_active"] == "true" {
//				// Call the next handler
//				c.Set("id", claim["id"])
//				c.Set("role_id", claim["role_id"])
//				c.Next()
//			} else {
//				c.JSON(http.StatusUnauthorized, dto.DefaultUnauthorizedResponse())
//				c.Abort()
//				return
//			}
//		}
//	}
//}

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

			if (claim["role_id"] == float64(1) || claim["role_id"] == float64(2)) && claim["is_verified"] == "true" && claim["is_active"] == "true" {
				// Call the next handler
				c.Set("id", claim["id"])
				c.Set("role_id", claim["role_id"])
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, dto.DefaultUnauthorizedResponse())
				c.Abort()
				return
			}
		}
	}
}
