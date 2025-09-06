package config

// Config armazena as configurações da aplicação.
// Por enquanto, está vazia, mas pode ser preenchida com:
// - String de conexão com o banco de dados
// - Chaves de API
// - Configurações de ambiente (dev, prod)
type Config struct {
	// Exemplo: Port string
}

// NewConfig carrega as configurações da aplicação.
func NewConfig() (*Config, error) {
	// A lógica para carregar de variáveis de ambiente ou arquivos .env iria aqui.
	return &Config{}, nil
}
