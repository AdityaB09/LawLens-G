package main

import (
	"log"
)

func main() {
	cfg := LoadConfig()
	db, err := ConnectDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	if err := AutoMigrate(db); err != nil {
		log.Fatalf("failed to migrate DB: %v", err)
	}

	router := SetupRouter(db)
	log.Printf("LawLens-G API listening on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
