package controllers

import (
	"api_gin/database"
	"api_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


// ExibeTodosAlunos godoc
// @Summary      Mostra todos alunos
// @Description  Busca a lista de todos os alunos
// @Tags         Alunos
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Router       /alunos [get]
func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)

	c.JSON(200, alunos)
}

// Saudacao godoc
// @Summary      Franse de boas vindas
// @Description  Retorna frase de boas vindas ao usuário que executou a chamada
// @Tags         Alunos
// @Accept       json
// @Produce      json
// @Param        nome path string true "nome do usuário"
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Router       /:nome [get]
func Saudacao(c *gin.Context){
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API diz:" : "E ai" + nome + ", tudo beleza?;",
	})
}

// CriaNovoAluno godoc
// @Summary      Adiciona aluno
// @Description  Adiciona um novo aluno no banco
// @Tags         Alunos
// @Accept       json
// @Produce      json
// @Param        aluno body models.Aluno true "id do aluno"
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Router       /alunos [post]
func CriaNovoAluno(c *gin.Context){
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if err := models.ValidaDadosAlunos(&aluno); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorID godoc
// @Summary      Busca aluno por ID
// @Description  Busca aluno por ID no banco
// @Tags         Alunos
// @Accept       json
// @Produce      json
// @Param        id   path string true "id do aluno"
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Router       /alunos/:id [get]
func BuscaAlunoPorID(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Aluno não encontrado"})
		return
	}
	
	c.JSON(http.StatusOK, aluno)		
}

// DeletaAluno godoc
// @Summary      Deleta aluno por ID
// @Description  Apaga aluno po ID
// @Tags         Alunos
// @Accept       json
// @Produce      json
// @Param        id   path string true "id do aluno"
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Router       /alunos/:id [delete]
func DeletaAluno(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data":"Aluno apagado com sucesso"})	
}

// EditaAluno godoc
// @Summary      Edita aluno
// @Description  Edita informações do alguno já cadastrado
// @Tags         Alunos
// @Accept       json
// @Produce      json
// @Param        id   path string true "id do aluno"
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Router       /alunos/:id [patch]
func EditaAluno(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})	
		return
	}
	if err := models.ValidaDadosAlunos(&aluno); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoCPF godoc
// @Summary      Busca aluno por CPF
// @Description  Busca aluno pelo CPF
// @Tags         Alunos
// @Accept       json
// @Produce      json
// @Param        cpf  path string true "cpf do aluno"
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Router       /alunos/cpf/:cpf [patch]
func BuscaAlunoCPF(c *gin.Context){
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}