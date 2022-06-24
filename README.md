# Apresentação

Programa exemplo desenvolvido em Go Lang.
É uma API desenvolvida com framework GIN, para persistir dados de cadastro de alunos com seus nomes, cpf e rg.


# Requisitos

Para executar o programa é necessário:
* Ter instalado o sdk do Go no seu sistema operacional. Pode ser pego em [https://go.dev/](https://go.dev/) 


# Recursos abordados

* Funções.
* Orientação de objetos.
* Modularização de packages.
* Utilização de rotas.
* MVC.
* Comunicação com banco de dados AWS em Mysql.
* CRUD completo em banco.
* Utilização de pacotes de terceiros.
* Validação de dados
* Testes unitários.


# Pacotes adicionais

* Para criação da API e gerenciamento do seu funcionamento: https://github.com/gin-gonic/gin, instalar go get -u github.com/gin-gonic/gin
* Para rotas: https://github.com/gorilla/mux
* Para ORM para banco : https://gorm.io/index.html, para instalar "go get -u gorm.io/gorm"
* Driver Mysql para ORM: "go get gorm.io/driver/mysql"
* Para validação de dados: https://pkg.go.dev/gopkg.in/validator.v2, instalar "go get gopkg.in/validator.v2"
* Para testes: https://github.com/stretchr/testify, instalar "go get github.com/stretchr/testify"


# Compilar e roda

Para compilar        : go build main.go
Para compilar e rodar: go run main.go

Para corrigir problemas de referencia de packages ou para possibilitar rodar packages fora do GO_HOME:

"go mod init"
