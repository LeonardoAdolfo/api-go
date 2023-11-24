package controllers

import (
	"api/initializers"
	model "api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostCreate lida com a criação de um novo post.
func PostCreate(c *gin.Context) {

	// Define uma struct para vincular os parâmetros do corpo da requisição.
	var body struct {
		Body  string
		Title string
	}

	// Vincula o corpo da requisição à struct definida.
	c.Bind(&body)

	// Cria uma nova instância de post com o título e corpo fornecidos.
	post := model.Post{Title: body.Title, Body: body.Body}

	// Tenta criar o post no banco de dados.
	result := initializers.DB.Create(&post)

	// Verifica se houve um erro durante o processo de criação.
	if result.Error != nil {
		// Se houver um erro, retorna um status 400 Bad Request.
		c.Status(400)
		return
	}

	// Se bem-sucedido, retorna o post criado na resposta.
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// PostIndex recupera todos os posts do banco de dados e os retorna na resposta.
func PostIndex(c *gin.Context) {
	var posts []model.Post

	// Recupera todos os posts do banco de dados.
	initializers.DB.Find(&posts)

	// Retorna os posts recuperados na resposta.
	c.JSON(http.StatusOK, gin.H{
		"post": posts,
	})
}

// PostShow recupera um post específico pelo ID do banco de dados e o retorna na resposta.
func PostShow(c *gin.Context) {
	id := c.Param("id")
	var post model.Post

	// Recupera o post com o ID especificado do banco de dados.
	initializers.DB.First(&post, id)

	// Retorna o post recuperado na resposta.
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// PostUpdate atualiza um post específico pelo ID com os dados fornecidos no corpo da requisição.
func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	// Define uma struct para vincular os parâmetros do corpo da requisição.
	var body struct {
		Body  string
		Title string
	}

	// Vincula o corpo da requisição à struct definida.
	c.Bind(&body)

	var post model.Post

	// Recupera o post com o ID especificado do banco de dados.
	initializers.DB.First(&post, id)

	// Atualiza o título e o corpo do post com os dados fornecidos.
	initializers.DB.Model(&post).Updates(model.Post{Title: body.Title, Body: body.Body})

	// Retorna o post atualizado na resposta.
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// PostDelete exclui um post específico pelo ID do banco de dados.
func PostDelete(c *gin.Context) {
	id := c.Param("id")

	// Exclui o post com o ID especificado do banco de dados.
	initializers.DB.Delete(&model.Post{}, id)

	// Retorna um status 200 OK para indicar exclusão bem-sucedida.
	c.Status(200)
}
