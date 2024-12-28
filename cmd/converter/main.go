package main

import (
	"converter/internal/config"
	"converter/internal/controllers/rest/router"
	"converter/internal/logs"
	"converter/internal/modules"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	logger, err := logs.InitLogger(cfg.Env, cfg.Log.LogFile)
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer logger.Close()

	l := slog.Default()

	handlers := modules.RegisterModules(l)

	r := gin.Default()

	r.Static("/", "./src")

	router.RegisterRoutes(r, handlers)

	port := fmt.Sprintf("%d", cfg.Server.Port)

	if err := r.Run(fmt.Sprintf(":" + port)); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
