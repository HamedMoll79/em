package main

import (
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"gitlab.sazito.com/sazito/event_publisher/adapter/redisqueue"
	"gitlab.sazito.com/sazito/event_publisher/config"
	"gitlab.sazito.com/sazito/event_publisher/pkg/postgresql"
	"gitlab.sazito.com/sazito/event_publisher/repository/migrator"
	"log"
)

var migrateFlag = flag.String("migrate", "", "Run migration up or down")

func main() {
	fmt.Println("\n\n\n\n\n start \n\n\n\n\n")
	//cfg, err := config.Load()
	//if err != nil {
	//	panic(err)
	//}

	cfg := config.Config{
		HTTPServer: config.HTTPServer{},
		Redis:      redisqueue.Config{},
		Postgres: postgresql.Config{
			Username: "sazito",
			Password: "Sazito123",
			Port:     5432,
			Host:     "localhost",
			DBName:   "sazito_evnet_manager",
			Driver:   "postgres",
			Schema:   "",
		},
	}

	controller := postgresql.NewPgController(postgresql.Config{
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		Port:     cfg.Postgres.Port,
		Host:     cfg.Postgres.Host,
		DBName:   cfg.Postgres.DBName,
		Driver:   cfg.Postgres.Driver,
		Schema:   cfg.Postgres.Schema,
	}, false)
	err := controller.Generate()
	fmt.Println("\n after generate \n")
	if err != nil {
		log.Println("Error controller.Generate", err)
		panic(err)
	}

	fmt.Println("\n before init \n")
	err = controller.Init()
	if err != nil {
		log.Println("Error controller.Init", err)
		panic(err)
	}

	mgr := migrator.New(controller.GetDataContext(), "./repository/postgresql/migrations")
	migrateOperation(*migrateFlag, mgr)
}

func migrateOperation(flag string, mg migrator.Migrator) {
	fmt.Printf("\n\n\n\n\n flag : %s \n\n\n\n\n", flag)
	switch flag {
	case "up":
		mg.Up()
	case "down":
		mg.Down()
	case "":
	default:
		log.Println("flag value is invalid, this flag only accepts the following values: up, down.")
	}
}
