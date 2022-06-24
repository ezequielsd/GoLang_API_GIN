// Desenvolvimento de api Rest usando o framework Gin:
// https://github.com/gin-gonic/gin, instalar go get -u github.com/gin-gonic/gin
// ORM para banco: https://gorm.io/index.html, para instalar "go get -u gorm.io/gorm", e driver mysql "go get gorm.io/driver/mysql" , para instalar driver do banco"
// Para validação de dados: https://pkg.go.dev/gopkg.in/validator.v2, instalar "go get gopkg.in/validator.v2"
// Para testes: https://github.com/stretchr/testify, instalar "go get github.com/stretchr/testify"
// Documentação: Gin Swagger https://github.com/swaggo/gin-swagger, instalar ""

package main

import (
	"api_gin/database"	
	"api_gin/routes"
)


func main() {
	database.ConectaDataBase()
	
	routes.HandleRequests()
}

