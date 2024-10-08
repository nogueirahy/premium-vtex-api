package internal

import (
	"log"
	"os"
	"prime/internal/prime"
	"prime/pkg/vtex"

	"github.com/joho/godotenv"
)

func getBaseURL() string {
	return "https://" + os.Getenv("ACCOUNT_NAME") + "." + os.Getenv("ENVIRONMENT") + ".com.br/api"
}

type AppDependencies struct {
	prime.PrimeHandler
}

func Setup() AppDependencies {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	baseURL := getBaseURL()

	vtexClient := vtex.NewVtexClient(baseURL)
	primeRepository := prime.NewPrimeRepository(vtexClient)
	primeService := prime.NewPrimeService(primeRepository)
	primeHandler := prime.NewPrimeHandler(primeService)

	return AppDependencies{
		PrimeHandler: primeHandler,
	}
}
