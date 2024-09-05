package services


import (
	"RankEdge/internal/models"
	"RankEdge/internal/repositories"
	"RankEdge/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

// UserService provides user services
type UserService interface {
	SignUp(user *models.User) error
	SignIn(email, password string) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}

// SignUp creates a new user
func (s *userService) SignUp(user *models.User) error {
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create the user
	user.Password = string(hash)
	user.UserID = utils.GenerateRandomID()

	if err := s.userRepo.CreateUser(user); err != nil {
		return err
	}

	return nil
}

// SignIn signs in a user
func (s *userService) SignIn(email, password string) (*models.User, error) {
	user, err := s.userRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Compare the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}


