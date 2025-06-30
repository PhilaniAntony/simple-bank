package api

import (
	"net/http"
	"strings"

	"github.com/PhilaniAntony/simplebank/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Message": "Authorization header not provided"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Verify the token
		fields := strings.Fields(authHeader)
		if len(fields) < 2 || fields[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Message": "Invalid authorization header format"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Message": "Authorization type must be Bearer"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := fields[1]
		payload, err := tokenMaker.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Message": "Invalid or expired token"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Store the payload in the context for use in handlers
		c.Set(authorizationPayloadKey, payload)
		c.Next()
	}
}
