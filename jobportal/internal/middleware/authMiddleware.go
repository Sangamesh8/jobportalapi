package middleware

import (
	"context"
	"errors"
	"job-portal-api/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// * this for
func (m *Mid) Authenticate(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		traceId, ok := ctx.Value(TraceIdKey).(string)
		if !ok {
			log.Error().Msg("tracker Id not Found!")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
			return
		}
		authHeader := c.Request.Header.Get("Authorization")
		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			err := errors.New("expected authorization header with: Bearer")
			log.Error().Err(err).Str("Trace Id", traceId).Msg("token not with bearer")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": http.StatusUnauthorized})
			return
		}
		claims, err := m.auth.ValidateToken(tokenParts[1])
		if err != nil {
			log.Error().Err(err).Str("Trace Id", traceId).Send()
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
			return
		}
		ctx = context.WithValue(ctx, auth.Authkey, claims)
		req := c.Request.WithContext(ctx)
		c.Request = req

		next(c)
	}
}
