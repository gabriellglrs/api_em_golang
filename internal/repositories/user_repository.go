package repositories

import (
	"database/sql"
	"fmt"

	"api-em-go/internal/models"
)

// UserRepository define a interface para o acesso a dados de usuários.
// Isso permite que a camada de serviço seja desacoplada da implementação do banco de dados.
type UserRepository interface {
	Create(user *models.User) error
	GetByID(id int64) (*models.User, error)
}

// postgresUserRepository é a implementação do UserRepository para o PostgreSQL.
type postgresUserRepository struct {
	db *sql.DB
}

// NewPostgresUserRepository cria uma nova instância do repositório de usuário para Postgres.
func NewPostgresUserRepository(db *sql.DB) UserRepository {
	return &postgresUserRepository{db: db}
}

// Create insere um novo usuário no banco de dados.
func (r *postgresUserRepository) Create(user *models.User) error {
	stmt, err := r.db.Prepare("INSERT INTO users(name, email) VALUES($1, $2) RETURNING id")
	if err != nil {
		return fmt.Errorf("erro ao preparar statement de inserção: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("erro ao executar inserção: %w", err)
	}
	return nil
}

// GetByID busca um usuário pelo ID no banco de dados.
func (r *postgresUserRepository) GetByID(id int64) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Usuário não encontrado
		}
		return nil, fmt.Errorf("erro ao buscar usuário por ID: %w", err)
	}
	return user, nil
}
