package database

import (
	"api_gin/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConectaDataBase(){
	//string de conexão para banco na AWS. Retirado na documentação, basta apenas trocar os campos de user, psw e endpoint.
	dsn := "user:psw@tcp(endpoint do db na AWS:3306)/api?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar no banco de dados")
	}else{
		log.Println("Conexão realizado com sucesso")
	}
	DB.AutoMigrate(&models.Aluno{})
}