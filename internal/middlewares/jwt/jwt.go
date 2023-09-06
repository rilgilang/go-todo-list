package jwt // Service is an interface from which our api module can access our repository of all our models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"simple-todo-list/internal/api/presenter"
	"simple-todo-list/internal/consts"
	"simple-todo-list/internal/entities"
	"simple-todo-list/internal/repositries"
	"strings"
	"time"
)

type AuthMiddleware interface {
	GenerateToken(user *entities.User, expirationMinute int, jwtKey string) (*string, error)
	ValidateToken() fiber.Handler
}

type authMiddlewares struct {
	userRepo repositries.UserRepository
}

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewAuthMiddleware(userRepo repositries.UserRepository) AuthMiddleware {
	return &authMiddlewares{
		userRepo: userRepo,
	}
}

func (m *authMiddlewares) GenerateToken(user *entities.User, expirationMinute int, jwtKey string) (*string, error) {
	// Declare the expiration time of the token
	expirationTime := time.Now().Add(time.Duration(expirationMinute) * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return nil, err
	}

	return &tokenString, nil
}

func (m *authMiddlewares) ValidateToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := strings.Split(c.GetReqHeaders()["Authorization"], "Bearer ")

		if len(authorization) != 2 {
			c.Status(400)
			return c.JSON(presenter.AuthErrorResponse(errors.New("token not valid!")))

		}

		token := authorization[1]

		// Initialize a new instance of `Claims`
		claims := &Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("my_secret_key"), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Status(401)
				return c.JSON(presenter.AuthErrorResponse(err))
			}
			c.Status(400)
			return c.JSON(presenter.AuthErrorResponse(err))
		}
		if !tkn.Valid {
			c.Status(401)
			return c.JSON(presenter.AuthErrorResponse(err))
		}

		c.Locals(consts.UserId, claims.ID)
		return c.Next()
	}
}
