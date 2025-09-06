package models

// User representa um usuário no sistema.
// As tags `json:"..."` são usadas para mapear os campos para JSON.
type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
