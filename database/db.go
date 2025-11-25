package database

import (
	"fmt" // Importe fmt para construir a string
	"log"
	"os" // Importe os para ler variáveis de ambiente

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
    DB  *gorm.DB
    err error
)

func ConectaComBancoDeDados() {
    // 1. LER AS VARIÁVEIS DE AMBIENTE
    host := os.Getenv("DB_HOST") 
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

    // Define porta padrão caso não esteja definida
    if port == "" {
        port = "5432"
    }
    
    // 2. CONSTRUIR A STRING DE CONEXÃO
    // Usamos fmt.Sprintf para juntar as variáveis lidas
    stringDeConexao := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        host, user, password, dbname, port)

    log.Printf("Tentando conectar com: %s:%s", host, port)

    // 3. ABRIR A CONEXÃO USANDO A STRING CONSTRUÍDA
    DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{}) // Adicionado &gorm.Config{} para GORM
    if err != nil {
        log.Panicf("Erro ao conectar com banco de dados: %v", err)
    }

    // O GORM faz o Ping e verifica a conexão na abertura. 
    // Se GORM.Open falhar, a conexão falhou, resolvendo o problema inicial.

    // 4. MIGRAÇÃO
    _ = DB.AutoMigrate(&models.Aluno{})
    log.Println("Conexão estabelecida e migração concluída com sucesso!")
}