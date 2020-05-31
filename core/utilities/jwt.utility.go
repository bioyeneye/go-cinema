package utilities

import (
	"fmt"
	"github.com/bioyeneye/rest-gin-api/core/configs"
	"github.com/bioyeneye/rest-gin-api/core/middlewares"
	"github.com/gin-gonic/gin"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AccessTokenResponse struct {
	AccessToken  string
	TokenType    string
	ExpiresIn    int64
	RefreshToken string
}

type IJWTService interface {
	GenerateToken(name string, admin bool) (AccessTokenResponse, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type JWTService struct {
	secretKey                    string
	issuer                       string
	refreshTokenExpiresInMinutes int64
	tokenExpiresInMinutes        int64
}

func NewJWTService() IJWTService {
	return &JWTService{
		secretKey:                    GetEnv(configs.JWTSecretKey, "secret"),
		issuer:                       "bolaji.com",
		refreshTokenExpiresInMinutes: GetEnvAsInt64(configs.JWTRefereshTokenExpiresInKey, 2),
		tokenExpiresInMinutes:        GetEnvAsInt64(configs.JWTTokenExpiresInKey, 1),
	}
}

func (jwtSrv *JWTService) GenerateToken(username string, admin bool) (AccessTokenResponse, error) {

	expiresAt := time.Now().Add(time.Hour * time.Duration(jwtSrv.tokenExpiresInMinutes)).Unix()
	claims := &jwtCustomClaims{
		username,
		admin,
		jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: expiresAt,
			Id:        username,
			IssuedAt:  time.Now().Unix(),
			Issuer:    jwtSrv.issuer,
			NotBefore: 0,
			Subject:   username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		panic(err)
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * time.Duration(jwtSrv.refreshTokenExpiresInMinutes)).Unix()
	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}

	return AccessTokenResponse{
		AccessToken:  t,
		TokenType:    "bearer",
		ExpiresIn:    int64(time.Hour/time.Millisecond) * jwtSrv.tokenExpiresInMinutes,
		RefreshToken: rt,
	}, err
}

func (jwtSrv *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(jwtSrv.secretKey), nil
	})
}

func GetCurrentUser(ctx *gin.Context) (middlewares.CurrentUser, bool) {
	currentuser, exists := ctx.MustGet(configs.ApplicationUserKey).(middlewares.CurrentUser)
	return currentuser, exists
}
