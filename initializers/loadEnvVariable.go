package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvVariables carrega variáveis de ambiente de um arquivo .env.
func LoadEnvVariables() {
	// Tenta carregar as variáveis de ambiente do arquivo .env.
	err := godotenv.Load()

	// Verifica se ocorreu um erro durante o carregamento.
	if err != nil {
		// Se houver um erro, registra um erro fatal.
		log.Fatal("Erro ao carregar o arquivo .env:", err)
	}
}
