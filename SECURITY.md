# Política de Segurança da Informação - Projeto Go/Gin (API Alunos)

## 1. Introdução

Esta Política de Segurança da Informação (PSI) estabelece as diretrizes e os requisitos mínimos de segurança para o desenvolvimento, implantação e operação do projeto **API de Gerenciamento de Alunos** (projeto_go_alura), desenvolvido em Go (Golang) com o framework Gin e banco de dados PostgreSQL.

O objetivo é proteger a confidencialidade, integridade e disponibilidade dos dados de alunos e da infraestrutura da aplicação contra ameaças internas e externas.

## 2. Escopo

Esta política se aplica a:
*   **Código-fonte e Dependências:** Todo o código-fonte e as bibliotecas de terceiros utilizadas no projeto.
*   **Ambientes:** Ambientes de desenvolvimento, teste e produção da aplicação.
*   **Pessoal:** Todos os desenvolvedores, colaboradores e administradores de sistema envolvidos no projeto.
*   **Dados:** Dados de alunos armazenados no banco de dados PostgreSQL.

## 3. Segurança da Aplicação (Go e Gin)

### 3.1. Validação de Entrada de Dados

Todos os dados recebidos pela API (parâmetros de rota, *query strings* e corpo da requisição JSON) **DEVEM** ser validados.

O projeto utiliza a biblioteca `gopkg.in/validator.v2` e *struct tags* nos modelos (`models/alunos.go`) para impor regras de formato e presença (ex: `nonzero`, `len=9`, `regexp`). Esta é a **implementação obrigatória** para garantir a integridade dos dados.
Qualquer falha na validação **DEVE** resultar em uma resposta HTTP `400 Bad Request` com uma mensagem de erro clara, sem expor detalhes internos da aplicação.

### 3.2. Tratamento de Erros e Logs

Mensagens de erro retornadas ao cliente **NÃO DEVEM** conter informações sensíveis, como *stack traces*, detalhes de conexão com o banco de dados ou caminhos de arquivos.
Logs de erro e de acesso **DEVEM** ser configurados para registrar eventos importantes, mas **NÃO DEVEM** armazenar dados sensíveis (RG, CPF, senhas, etc.) em texto simples.

### 3.3. Gerenciamento de Dependências

Todas as dependências (módulos Go) **DEVEM** ser mantidas atualizadas para mitigar vulnerabilidades conhecidas.
O comando `go get -u all` e a verificação regular de vulnerabilidades (`go mod tidy` e ferramentas de análise estática) **DEVEM** ser executados antes de cada *release*.

## 4. Segurança de Dados (PostgreSQL e GORM)

### 4.1. Proteção de Dados Sensíveis

O projeto lida com dados sensíveis (RG e CPF).

Embora o projeto atual armazene RG e CPF em texto simples, para um ambiente de produção, estes campos **DEVEM** ser criptografados no banco de dados (criptografia em repouso) ou, no mínimo, armazenados usando funções de *hashing* seguras se a recuperação do valor original não for necessária.
A conexão com o banco de dados **DEVE** ser feita usando credenciais dedicadas e, em ambientes de produção, através de conexões criptografadas (SSL/TLS).

### 4.2. Credenciais de Acesso

Credenciais de banco de dados (usuário, senha) **NÃO DEVEM** ser codificadas diretamente no código-fonte.
O projeto utiliza **Variáveis de Ambiente** (`DB_HOST`, `DB_USER`, etc.), o que é a abordagem correta. Em produção, estas variáveis **DEVEM** ser gerenciadas por um serviço de gerenciamento de segredos (ex: HashiCorp Vault, AWS Secrets Manager).

### 4.3. Prevenção de Injeção SQL

Todas as operações de banco de dados **DEVEM** utilizar *prepared statements* ou o ORM (GORM) para evitar ataques de Injeção SQL.
O uso do GORM no projeto (`database/db.go`, `controllers/controller.go`) já mitiga este risco, mas consultas manuais **DEVEM** sempre usar *placeholders* e parâmetros.

## 5. Segurança de Infraestrutura (Docker)

### 5.1. Imagens Base Seguras

O `Dockerfile` **DEVE** utilizar imagens base oficiais e mínimas.
O uso de *multi-stage build* (fase `builder` com `golang:1.25` e fase final com `scratch`) é uma excelente prática que minimiza a superfície de ataque da imagem final.

### 5.2. Isolamento de Contêineres

O contêiner da aplicação **NÃO DEVE** rodar com privilégios de *root* (embora o `scratch` já ajude nisso, o usuário dentro do contêiner **DEVE** ser não-root).
O contêiner do PostgreSQL **DEVE** ter seus dados persistidos em um volume dedicado (`postgres_data`), e o acesso externo **DEVE** ser restrito apenas ao necessário (a porta `5432` só deve ser exposta para o host se for estritamente necessário para desenvolvimento/administração).

## 6. Desenvolvimento Seguro e Revisão de Código

### 6.1. Revisão de Código (Code Review)

Todas as alterações no código **DEVEM** ser revisadas por pelo menos um outro desenvolvedor antes de serem mescladas na *branch* principal (`main`).
A revisão **DEVE** incluir a verificação de vulnerabilidades comuns de API (OWASP Top 10), como falhas de autenticação, exposição de dados e validação inadequada.

### 6.2. Testes de Segurança

Testes de unidade e integração **DEVEM** ser escritos para cobrir a lógica de validação e as regras de negócio.
Adoção de testes de segurança automatizados (SAST - Static Application Security Testing) no pipeline de CI/CD para analisar o código Go em busca de vulnerabilidades.

## 7. Resposta a Incidentes

Em caso de suspeita ou confirmação de um incidente de segurança (ex: vazamento de dados, acesso não autorizado), a equipe **DEVE** seguir o seguinte procedimento:
    1. **Contenção:** Isolar o sistema afetado (ex: derrubar o contêiner da aplicação).
    2. **Análise:** Coletar logs e evidências para determinar a causa raiz.
    3. **Remediação:** Corrigir a vulnerabilidade e restaurar o serviço.
    4. **Comunicação:** Notificar as partes interessadas (incluindo o cliente/usuário, se aplicável) sobre o incidente e as medidas tomadas.

## 8. Referências

[1] [OWASP Top 10 - 2021](https://owasp.org/www-project-top-ten/)
[2] [Go Security Guide](https://go.dev/doc/security/guide)
[3] [Gin Gonic Security](https://gin-gonic.com/docs/introduction/security/)
[4] [GORM Security](https://gorm.io/docs/security.html)
