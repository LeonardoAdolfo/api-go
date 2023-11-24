package main

import (
	"api/controllers"
	"api/initializers"

	"github.com/gin-gonic/gin"
)

// init é uma função especial que é executada antes da função main e é usada para inicialização.
func init() {
	// Carrega as variáveis de ambiente do arquivo .env.
	initializers.LoadEnvVariables()

	// Estabelece uma conexão com o banco de dados.
	initializers.ConnectToDB()
}

func main() {
	// Cria uma instância do framework Gin.
	r := gin.Default()

	// Define rotas para manipular operações CRUD em posts.
	r.POST("/post", controllers.PostCreate)       // Rota para criar um novo post.
	r.GET("/post", controllers.PostIndex)         // Rota para obter todos os posts.
	r.GET("/post/:id", controllers.PostShow)      // Rota para obter um post específico por ID.
	r.PUT("/post/:id", controllers.PostUpdate)    // Rota para atualizar um post existente por ID.
	r.DELETE("/post/:id", controllers.PostDelete) // Rota para excluir um post por ID.

	// Inicia o servidor da aplicação.
	r.Run()
}
