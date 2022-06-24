package main

import (
	"api_gin/controllers"
	"api_gin/database"
	"api_gin/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"	
)

var ID int

func SetupRotasTestes() *gin.Engine{
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestListaTodosOsAlunosHanlder(t *testing.T) {
	database.ConectaDataBase()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasTestes()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBucaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaDataBase()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasTestes()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorIdHandler(t *testing.T){
	database.ConectaDataBase()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasTestes()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	pathBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	fmt.Println(alunoMock.Nome)
	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome )
}

func TestDeletaAlunoHandler(t *testing.T){
	database.ConectaDataBase()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasTestes()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	pathBusca := "/alunos/" + strconv.Itoa(ID)
	req, _:= http.NewRequest("DELETE", pathBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditAlunoHandler(t *testing.T){
	database.ConectaDataBase()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasTestes()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "47123456789", RG: "123456700"}
	valorJson, _:=json.Marshal(aluno)
	pathEditar := "/alunos/" + strconv.Itoa(ID)
	req, _:=http.NewRequest("PATCH", pathEditar, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "47123456789", alunoMockAtualizado.CPF)
	assert.Equal(t, "123456700", alunoMockAtualizado.RG)
	assert.Equal(t, "Nome do Aluno Teste", alunoMockAtualizado.Nome)
}