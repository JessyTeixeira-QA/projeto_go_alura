package routes

import (
	"net/http" // Importamos para usar http.StatusOK e outras constantes

	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	// --- 🚀 ROTAS ADICIONADAS PARA RESOLVER O 404 ---
	// 1. Rota de Health Check/Ping (retorna "pong")
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 2. Rota Raiz (Exibe uma mensagem simples de status da API)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "API Alunos está rodando!", "documentacao_html": "/index"})
	})
	// -----------------------------------------------

	// Rotas corrigidas para evitar ambiguidades:

	// Rota de Saudação: Movemos para uma URL específica.
	r.GET("/alunos/saudacao/:nome", controllers.Saudacoes)

	// Rotas de Alunos (sem alteração de ordem)
	r.GET("/alunos", controllers.TodosAlunos)
	r.GET("/alunos/:id", controllers.BuscarAlunoPorID)
	r.POST("/alunos", controllers.CriarNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	r.PATCH("/alunos/:id", controllers.EditarAluno)

	// Rota de Busca por CPF (já estava correta)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	// Rotas de View/Outras
	r.GET("/index", controllers.ExibePaginaIndex)

	// Rota que trata qualquer caminho não encontrado
	r.NoRoute(controllers.RotaNaoEncontrada)

	// Roda o servidor. O Gin vai pegar a porta 8080 por padrão ou a variável de ambiente PORT.
	_ = r.Run()
}
