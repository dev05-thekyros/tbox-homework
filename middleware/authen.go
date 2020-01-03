package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"tbox-homework/modules/user/usrmodel"
)

type UserAuthStorage interface {
	GetUserByAccessToken(ctx context.Context, accessToken string) (*usrmodel.User, error)
}
type authMiddleware struct {
	userStorage UserAuthStorage
}

func NewAuthMiddleware(userStorage UserAuthStorage) *authMiddleware {
	return &authMiddleware{userStorage}
}

func (am *authMiddleware) AuthRequired(c *gin.Context) {
	accessToken := c.GetHeader("access_token")
	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	usr, err := am.userStorage.GetUserByAccessToken(context.Background(), accessToken)
	if err != nil || usr == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Set("user", usr)
	c.Next()
}
