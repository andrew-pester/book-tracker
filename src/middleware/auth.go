package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	ResourceAccess Client `json:"resource_access,omitempty"`
	Jti            string `json:"jti,omitempty"`
}

type Client struct {
	BookTrackerClient ClientRoles `json:"book-tracker-client,omitempty"`
}
type ClientRoles struct {
	Roles []string `json:"roles,omitempty"`
}

func IsAuthorized(c *gin.Context) error {
	accessToken, err := getToken(c)
	if err != nil {
		return err
	}

	token, err := verifyToken(accessToken)
	if err != nil {
		return err
	}

	if err := checkClaims(token); err != nil {
		return err
	}

	return nil
}
func getToken(c *gin.Context) (string, error) {
	accessTokenRaw := c.GetHeader("Authorization")
	if len(accessTokenRaw) == 0 {
		return "", errors.New("Authorization header not given")
	}
	headerInfo := strings.Split(accessTokenRaw, " ")
	if len(headerInfo) < 2 {
		return "", errors.New("Header does not contain needed information")
	}
	return headerInfo[1], nil

}

func verifyToken(accessToken string) (*oidc.IDToken, error) {
	client := &http.Client{}
	ctx := oidc.ClientContext(context.Background(), client)

	provider, err := oidc.NewProvider(ctx, os.Getenv("PROVIDER_URL"))
	if err != nil {
		return nil, err
	}

	config := &oidc.Config{
		ClientID: os.Getenv("CLIENT_ID"),
	}

	verifier := provider.Verifier(config)
	token, err := verifier.Verify(ctx, accessToken)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func checkClaims(token *oidc.IDToken) error {
	var claim Claims
	if err := token.Claims(&claim); err != nil {
		return err
	}
	for _, i := range claim.ResourceAccess.BookTrackerClient.Roles {
		if i == os.Getenv("BOOK_ADMIN_ROLE") {
			return nil
		}
	}
	return errors.New("User does not have the required role")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := IsAuthorized(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
	}
}
