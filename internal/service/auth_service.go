package service

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"simple-todo-list/internal/consts"
	"simple-todo-list/internal/entities"
	"simple-todo-list/internal/middlewares/jwt"
	"simple-todo-list/internal/repositries"
)

// Service is an interface from which our api module can access our repository of all our models
type AuthService interface {
	Login(user *entities.User) (userData *entities.User, token *string, err error)
	Register(user *entities.User) (userData *entities.User, token *string, err error)
}

type authService struct {
	jwtMdwr  jwt.AuthMiddleware
	userRepo repositries.UserRepository
}

// NewService is used to create a single instance of the service
func NewAuthService(jwtMdwr jwt.AuthMiddleware, userRepo repositries.UserRepository) AuthService {
	return &authService{
		jwtMdwr:  jwtMdwr,
		userRepo: userRepo,
	}
}

func (s *authService) Login(user *entities.User) (*entities.User, *string, error) {
	userData, err := s.userRepo.FindOne(user.Username)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	if userData == nil {
		return nil, nil, errors.New("username or password is wrong!")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return nil, nil, errors.New("username or password is wrong!")
	}

	token, err := s.jwtMdwr.GenerateToken(userData, 20, "my_secret_key")

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	return userData, token, nil
}

func (s *authService) Register(user *entities.User) (*entities.User, *string, error) {

	checkUser, err := s.userRepo.FindOne(user.Username)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	if checkUser != nil {
		return nil, nil, errors.New("username already taken!")
	}

	saltByte, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	//set to salted password
	user.Password = string(saltByte)

	err = s.userRepo.CreateUser(user)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	newUser, err := s.userRepo.FindOne(user.Username)

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	token, err := s.jwtMdwr.GenerateToken(newUser, 5, "my_secret_key")

	if err != nil {
		return nil, nil, errors.New(consts.InternalServerError)
	}

	return newUser, token, nil
}
