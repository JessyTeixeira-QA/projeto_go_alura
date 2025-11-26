# 🎓 API de Gerenciamento de Alunos com Go e Gin (Alura)

Este repositório contém o projeto de uma **API RESTful** desenvolvida em **Go** (Golang) utilizando o framework **Gin** para o roteamento e o **GORM** como ORM para persistência de dados em um banco **PostgreSQL**.

O projeto foi estruturado para demonstrar a criação de uma aplicação completa, com foco em:
*   **Boas Práticas de Desenvolvimento:** Separação de responsabilidades (Controllers, Models, Database, Routes).
*   **Validação de Dados:** Uso de *struct tags* e a biblioteca `gopkg.in/validator.v2` para garantir a integridade dos dados.
*   **Contêineres:** Configuração completa com `Dockerfile` e `docker-compose.yml` para um ambiente de desenvolvimento e produção isolado.

## 🚀 Funcionalidades

A API permite o gerenciamento completo de registros de alunos (CRUD - Create, Read, Update, Delete).

| Rota | Método | Descrição | Controller |
| :--- | :--- | :--- | :--- |
| `/` | `GET` | Retorna o status da API. | N/A |
| `/ping` | `GET` | Health Check simples (retorna "pong"). | N/A |
| `/alunos` | `GET` | Lista todos os alunos cadastrados. | `TodosAlunos` |
| `/alunos` | `POST` | Cria um novo aluno. Requer validação de `Nome`, `RG` (9 dígitos) e `CPF` (11 dígitos). | `CriarNovoAluno` |
| `/alunos/:id` | `GET` | Busca um aluno pelo ID. | `BuscarAlunoPorID` |
| `/alunos/:id` | `PATCH` | Atualiza os dados de um aluno pelo ID. | `EditarAluno` |
| `/alunos/:id` | `DELETE` | Deleta um aluno pelo ID. | `DeletarAluno` |
| `/alunos/cpf/:cpf` | `GET` | Busca um aluno pelo número de CPF. | `BuscaAlunoPorCPF` |
| `/alunos/saudacao/:nome` | `GET` | Rota de exemplo que retorna uma saudação personalizada. | `Saudacoes` |
| `/index` | `GET` | Exibe uma página HTML simples com a lista de alunos (View). | `ExibePaginaIndex` |

## 🛠️ Tecnologias Utilizadas

*   **Linguagem:** Go (Golang)
*   **Framework Web:** [Gin Gonic](https://github.com/gin-gonic/gin)
*   **ORM:** [GORM](https://gorm.io/)
*   **Banco de Dados:** PostgreSQL
*   **Contêineres:** Docker e Docker Compose

## ⚙️ Configuração do Ambiente

O projeto utiliza **Docker Compose** para orquestrar a aplicação Go e o banco de dados PostgreSQL, facilitando a configuração do ambiente.

### Pré-requisitos

Certifique-se de ter o [Docker](https://www.docker.com/get-started) e o [Docker Compose](https://docs.docker.com/compose/install/) instalados em sua máquina.

### 1. Clonar o Repositório

```bash
git clone https://github.com/guilhermeonrails/api-go-gin.git
cd api-go-gin
```

### 2. Iniciar a Aplicação

Execute o comando abaixo para construir as imagens e iniciar os contêineres:

```bash
docker-compose up --build
```

O Docker Compose irá:
1.  Construir a imagem da aplicação Go (`app`) usando o `Dockerfile`.
2.  Iniciar o contêiner do PostgreSQL (`postgres`), aguardando que ele esteja saudável.
3.  Iniciar o contêiner da aplicação Go, que se conectará ao banco de dados e executará as migrações (criação da tabela `alunos`).

A aplicação estará acessível em `http://localhost:8080`.

### 3. Variáveis de Ambiente

A conexão com o banco de dados é configurada através de variáveis de ambiente definidas no `docker-compose.yml` e lidas pelo arquivo `database/db.go`.

| Variável | Valor Padrão | Descrição |
| :--- | :--- | :--- |
| `DB_HOST` | `postgres` | Nome do serviço do banco de dados no Docker Compose. |
| `DB_USER` | `root` | Usuário do PostgreSQL. |
| `DB_PASSWORD` | `root` | Senha do PostgreSQL. |
| `DB_NAME` | `root` | Nome do banco de dados. |
| `DB_PORT` | `5432` | Porta do PostgreSQL. |

## 💻 Como Rodar Localmente (Sem Docker)

Se preferir rodar a aplicação diretamente em sua máquina, siga os passos:

### Pré-requisitos

*   [Go (versão 1.24 ou superior)](https://golang.org/dl/)
*   Um servidor PostgreSQL rodando localmente.

### 1. Configurar o Banco de Dados

Crie um banco de dados PostgreSQL e configure as variáveis de ambiente necessárias para a conexão (substitua pelos seus dados):

```bash
export DB_HOST=localhost
export DB_USER=seu_usuario
export DB_PASSWORD=sua_senha
export DB_NAME=seu_banco
export DB_PORT=5432
```

### 2. Instalar Dependências e Rodar

```bash
go mod tidy
go run main.go
```

A aplicação será iniciada na porta `8080`.

## 📝 Exemplo de Uso da API (POST)

Para criar um novo aluno, envie uma requisição `POST` para a rota `/alunos` com um corpo JSON no formato:

```json
{
    "nome": "João da Silva",
    "rg": "123456789",
    "cpf": "12345678901"
}
```

**Exemplo com `curl`:**

```bash
curl -X POST http://localhost:8080/alunos \
-H "Content-Type: application/json" \
-d '{"nome": "João da Silva", "rg": "123456789", "cpf": "12345678901"}'
```
