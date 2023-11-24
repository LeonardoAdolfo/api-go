package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB é uma variável global que representa a conexão com o banco de dados.
var DB *gorm.DB

// ConnectToDB estabelece uma conexão com o banco de dados usando as configurações fornecidas no ambiente.
func ConnectToDB() {
	// Inicializa a variável de erro.
	var err error

	// Obtém a string de conexão do banco de dados do ambiente.
	dsn := os.Getenv("DB_URL")

	// Abre uma conexão com o banco de dados PostgreSQL usando a string de conexão.
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Verifica se ocorreu um erro durante a conexão.
	if err != nil {
		// Se houver um erro, registra um erro fatal.
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
}
