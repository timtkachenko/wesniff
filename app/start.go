package app

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"sync"
	controllers "wesniff/controller"
)

var (
	OnceInit    = &sync.Once{}
	userHandler *controllers.UserHandler
)

func config() {
	godotenv.Overload(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg("Error loading .env file. Using default config...")
	}
}
func Start() {
	OnceInit.Do(func() {
		config()
		userHandler = controllers.NewUserHandler()

	})
}

func Handler() *controllers.UserHandler {
	return userHandler
}
