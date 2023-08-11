package main

import (
	"github.com/TheNeoCarvalho/gopportunities/config"
	"github.com/TheNeoCarvalho/gopportunities/router"
)

var (
	logger *config.Logger
)
func main() {
	logger = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.Init()
}