package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

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
	// A rota r.GET("/alunos/", controllers.BuscaAlunoPorCPF) estava redundante, 
	// pois "/alunos" já chama TodosAlunos. Eu a removi.
	
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.Run()
}