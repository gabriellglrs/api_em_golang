package dto

// CreateUserRequest é o DTO para criar um novo usuário.
// Contém as informações necessárias que vêm no corpo da requisição.
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserResponse é o DTO para retornar informações de um usuário.
// Usado para evitar expor a estrutura completa do modelo do banco de dados.
type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
