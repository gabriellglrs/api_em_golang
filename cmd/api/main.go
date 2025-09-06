package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"api-em-go/internal/handlers"
	"api-em-go/internal/repositories"
	"api-em-go/internal/router"
	"api-em-go/internal/services"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

func main() {
	// Configuração da conexão com o banco de dados
	dbConnStr := "postgres://admin:admin@localhost:5432/api_db?sslmode=disable"
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
	}
	defer db.Close()

	// Testa a conexão com o banco de dados
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	fmt.Println("Conexão com o banco de dados PostgreSQL estabelecida com sucesso!")

	// Código temporário para executar migração
	migrationSQL, err := ioutil.ReadFile("migrations/001_create_users_table.sql")
	if err != nil {
		log.Fatalf("Erro ao ler arquivo de migração: %v", err)
	}
	_, err = db.Exec(string(migrationSQL))
	if err != nil {
		log.Fatalf("Erro ao executar migração: %v", err)
	}
	fmt.Println("Migração 001_create_users_table.sql executada com sucesso!")

	// Código temporário para executar migração de população
	migrationSQL2, err := ioutil.ReadFile("migrations/002_populate_users_table.sql")
	if err != nil {
		log.Fatalf("Erro ao ler arquivo de migração 002: %v", err)
	}
	_, err = db.Exec(string(migrationSQL2))
	if err != nil {
		log.Fatalf("Erro ao executar migração 002: %v", err)
	}
	fmt.Println("Migração 002_populate_users_table.sql executada com sucesso!")
	// Fim do código temporário de migração de população

	// Fim do código temporário

	// Injeção de Dependências
	userRepo := repositories.NewPostgresUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Cria um novo roteador e passa o userHandler
	r := router.NewRouter(userHandler)

	fmt.Println("Iniciando o servidor na porta 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
