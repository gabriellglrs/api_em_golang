# API em Go 🚀

Uma API RESTful simples desenvolvida em Go para fins de aprendizado, demonstrando conceitos fundamentais de desenvolvimento backend com Clean Architecture.

## 📚 Objetivos de Aprendizado

Este projeto foi criado para ensinar e demonstrar:

- **Clean Architecture** - Separação clara de responsabilidades
- **Dependency Injection** - Desacoplamento entre camadas
- **Repository Pattern** - Abstração da camada de dados
- **REST API** - Desenvolvimento de endpoints HTTP
- **PostgreSQL** - Integração com banco de dados relacional
- **Docker** - Containerização de aplicações

## 🏗️ Arquitetura do Projeto

```
api-em-go/
├── cmd/api/              # Ponto de entrada da aplicação
├── internal/
│   ├── config/          # Configurações da aplicação
│   ├── dto/             # Data Transfer Objects
│   ├── handlers/        # Controllers/Handlers HTTP
│   ├── models/          # Modelos de dados
│   ├── repositories/    # Camada de acesso a dados
│   ├── router/          # Configuração de rotas
│   └── services/        # Lógica de negócios
├── migrations/          # Scripts SQL de migração
├── docker-compose.yml   # Configuração do PostgreSQL
├── go.mod              # Gerenciamento de dependências
└── go.sum
```

### 🧱 Camadas da Arquitetura

1. **Handlers** (`internal/handlers/`): Recebem requisições HTTP e retornam respostas
2. **Services** (`internal/services/`): Contêm a lógica de negócios
3. **Repositories** (`internal/repositories/`): Fazem o acesso aos dados
4. **Models** (`internal/models/`): Definem as estruturas de dados
5. **DTOs** (`internal/dto/`): Objetos para transferência de dados

## 🚀 Como Executar

### Pré-requisitos

- Go 1.25+ instalado
- Docker e Docker Compose

### 1. Clone o repositório

```bash
git clone <seu-repositorio>
cd api-em-go
```

### 2. Inicie o banco de dados

```bash
docker-compose up -d
```

### 3. Instale as dependências

```bash
go mod tidy
```

### 4. Execute a aplicação

```bash
go run cmd/api/main.go
```

A API estará disponível em `http://localhost:8080`

## 📖 Endpoints da API

### Health Check
```http
GET /health
```

**Resposta:**
```json
{
  "status": "ok"
}
```

### Criar Usuário
```http
POST /users
Content-Type: application/json

{
  "name": "João Silva",
  "email": "joao@example.com"
}
```

**Resposta:**
```json
{
  "id": 1,
  "name": "João Silva",
  "email": "joao@example.com"
}
```

### Buscar Usuário por ID
```http
GET /users/{id}
```

**Resposta:**
```json
{
  "id": 1,
  "name": "João Silva",
  "email": "joao@example.com"
}
```

## 🧪 Testando a API

### Usando cURL

**Health Check:**
```bash
curl http://localhost:8080/health
```

**Criar usuário:**
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Maria Santos", "email": "maria@example.com"}'
```

**Buscar usuário:**
```bash
curl http://localhost:8080/users/1
```

## 🗄️ Banco de Dados

O projeto utiliza PostgreSQL com as seguintes configurações:

- **Host:** localhost:5432
- **Usuário:** admin
- **Senha:** admin
- **Database:** api_db

### Estrutura da Tabela Users

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);
```

## 🔧 Tecnologias Utilizadas

- **[Go](https://golang.org/)** - Linguagem de programação
- **[Gorilla Mux](https://github.com/gorilla/mux)** - Roteador HTTP
- **[PostgreSQL](https://www.postgresql.org/)** - Banco de dados
- **[lib/pq](https://github.com/lib/pq)** - Driver PostgreSQL para Go
- **[Docker](https://www.docker.com/)** - Containerização

## 🎯 Conceitos Demonstrados

### 1. Clean Architecture
A aplicação segue os princípios da Clean Architecture, mantendo as dependências apontando para dentro e separando claramente as responsabilidades.

### 2. Dependency Injection
As dependências são injetadas através dos construtores, permitindo fácil substituição e testabilidade.

### 3. Repository Pattern
A interface `UserRepository` abstrai o acesso aos dados, permitindo diferentes implementações (PostgreSQL, MongoDB, etc.).

### 4. DTOs (Data Transfer Objects)
Utilizamos DTOs para controlar exatamente quais dados são enviados e recebidos pela API.

## 🔄 Migrações

O projeto inclui migrações automáticas que são executadas na inicialização:

1. `001_create_users_table.sql` - Cria a tabela de usuários
2. `002_populate_users_table.sql` - Insere dados de teste

> **Nota:** Em produção, seria recomendado usar uma ferramenta de migração mais robusta como [migrate](https://github.com/golang-migrate/migrate).

## 📈 Próximos Passos

Para expandir seus conhecimentos, considere implementar:

- [ ] Validação de dados de entrada
- [ ] Middleware de logging
- [ ] Autenticação JWT
- [ ] Testes unitários e de integração
- [ ] Documentação Swagger/OpenAPI
- [ ] Paginação para listagem de usuários
- [ ] Soft delete
- [ ] Cache com Redis
- [ ] Containerização da aplicação Go

## 🤝 Contribuindo

Este é um projeto educacional! Sinta-se à vontade para:

1. Fazer fork do projeto
2. Criar uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abrir um Pull Request

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 📚 Recursos Adicionais

- [Documentação oficial do Go](https://golang.org/doc/)
- [Clean Architecture em Go](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [REST API Design Best Practices](https://restfulapi.net/)
- [PostgreSQL Tutorial](https://www.postgresql.org/docs/current/tutorial.html)

---

**Feito com ❤️ para aprendizado em Go**
