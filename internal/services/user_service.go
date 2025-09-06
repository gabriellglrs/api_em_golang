package services

import (
	"api-em-go/internal/models"
	"api-em-go/internal/repositories"
)

// UserService lida com a lógica de negócios para usuários.
type UserService struct {
	repo repositories.UserRepository
}

// NewUserService cria uma nova instância de UserService.
func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser cria um novo usuário usando o repositório.
func (s *UserService) CreateUser(user *models.User) error {
	// Aqui você pode adicionar lógica de validação de negócios antes de salvar.
	return s.repo.Create(user)
}

// GetUserByID busca um usuário pelo ID usando o repositório.
func (s *UserService) GetUserByID(id int64) (*models.User, error) {
	return s.repo.GetByID(id)
}
