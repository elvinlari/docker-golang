package middleware

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
	"os"

	"github.com/gin-gonic/gin"

	oidc "github.com/coreos/go-oidc"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json"
)

type Res401Struct struct {
	Status   string `json:"status" example:"FAILED"`
	HTTPCode int    `json:"httpCode" example:"401"`
	Message  string `json:"message" example:"authorisation failed"`
}

//claims component of jwt contains mainy fields , we need only roles of DemoServiceClient
//"DemoServiceClient":{"DemoServiceClient":{"roles":["pets-admin","pet-details","pets-search"]}},
type Claims struct {
	ResourceAccess client `json:"resource_access,omitempty"`
	JTI            string `json:"jti,omitempty"`
}

type client struct {
	DemoServiceClient clientRoles `json:"DemoServiceClient,omitempty"`
}

type clientRoles struct {
	Roles []string `json:"roles,omitempty"`
}

var RealmConfigURL string = os.Getenv("REALMCONFIG_URL")
// var RealmConfigURL string = "http://localhost:8080/realms/lynq"

var clientID string = os.Getenv("KEYCLOAK_CLIENTID") 


// CommonHeaders to share between packages
func CommonHeaders(h gin.HandlerFunc) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header(contentType, applicationJSON)
        h(c)
    }
}

// keycloak
func IsAuthorizedJWT(h gin.HandlerFunc, role string) gin.HandlerFunc {
    return func(c *gin.Context) {
        rawAccessToken := c.GetHeader("Authorization")

        tr := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }
        client := &http.Client{
            Timeout:   time.Duration(6000) * time.Second,
            Transport: tr,
        }
        ctx := oidc.ClientContext(context.Background(), client)
        provider, err := oidc.NewProvider(ctx, RealmConfigURL)
        if err != nil {
            authorisationFailed("authorisation failed while getting the provider: "+RealmConfigURL+err.Error(), c)
            return
        }

        oidcConfig := &oidc.Config{
            ClientID: clientID,
        }
        verifier := provider.Verifier(oidcConfig)
        idToken, err := verifier.Verify(ctx, rawAccessToken)
        if err != nil {
            authorisationFailed("authorisation failed while verifying the token: "+err.Error(), c)
            return
        }

        var IDTokenClaims Claims // ID Token payload is just JSON.
        if err := idToken.Claims(&IDTokenClaims); err != nil {
            authorisationFailed("claims : "+err.Error(), c)
            return
        }
        fmt.Println(IDTokenClaims)
        //checking the roles
        userAccessRoles := IDTokenClaims.ResourceAccess.DemoServiceClient.Roles
        for _, b := range userAccessRoles {
            if b == role {
                c.Header(contentType, applicationJSON)
                h(c)
                return
            }
        }

        authorisationFailed("user not allowed to access this api", c)
    }
}


func authorisationFailed(message string, c *gin.Context) {
    c.Header("Content-Type", "application/json; charset=utf-8")
    c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
        "Status":   "FAILED",
        "HTTPCode": http.StatusUnauthorized,
        "Message":  message,
    })
}