package main

import (
	"context"
	"flag"
	"fmt"
	"gitlab.sazito.com/sazito/event_publisher/adapter/redis_adapter"
	"gitlab.sazito.com/sazito/event_publisher/config"
	"gitlab.sazito.com/sazito/event_publisher/pkg/postgresql"
	"gitlab.sazito.com/sazito/event_publisher/repository/migrator"
	postgresql2 "gitlab.sazito.com/sazito/event_publisher/repository/postgresql"
	"gitlab.sazito.com/sazito/event_publisher/repository/redis_db"
	"gitlab.sazito.com/sazito/event_publisher/services/event_service"
	"log"
)

var migrateFlag = flag.String("migrate", "", "Run migration up or down")

func main() {
	flag.Parse()
	//cfg, err := config.Load()
	//if err != nil {
	//	panic(err)
	//}

	cfg := config.Config{
		HTTPServer: config.HTTPServer{},
		Redis:      redis_adapter.Config{},
		Postgres: postgresql.Config{
			Username: "sazito",
			Password: "Sazito123",
			Port:     5432,
			Host:     "em-database",
			DBName:   "sazito_event_manager",
			Driver:   "postgres",
			Schema:   "",
		},
	}

	/*pgdatabase := postgresql.PgDatabase{
		DBConfig: postgresql.Config{
			Username: "sazito",
			Password: "Sazito123",
			Port:     5432,
			Host:     "em-database",
			DBName:   "sazito_event_manager",
			Driver:   "postgres",
			Schema:   "",
		},
		Database: nil,
	}*/

	//fmt.Printf("pgdatabase: %+v\n", pgdatabase)

	//db, err := sql.Open("postgres", pgdatabase.DBConfig.DSN())

	//fmt.Printf("\nopen db error : %v \n ", err)

	//err = db.Ping()

	//fmt.Printf("\nping db error : %v \n ", err)

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
	if err != nil {
		panic(err)
	}

	err = controller.Init()
	if err != nil {
		panic(err)
	}

	mgr := migrator.New(controller.GetDataContext(), "./repository/postgresql/migrations")
	migrateOperation(*migrateFlag, mgr)

	services := setUpServices(cfg, controller)

	redisDB := services.redisDB

	pingResult, err := redisDB.Conn().Ping(context.Background()).Result()

	if err != nil {
		fmt.Printf("cant ping redis %v", err)
	}

	fmt.Printf("ping result %v\n", pingResult)
	//todo - redis connection
	//TODO- setup services
	//todo - setup router

	//todo - test fake publisher redis and consume
}

type GlobalServices struct {
	eventService event_service.Service
	redisDB      *redis_db.DB
}

func setUpServices(cfg config.Config, controller *postgresql.PgController) GlobalServices {
	rediscfg := redis_adapter.Config{
		UserName: "",
		Host:     "em-redis",
		Port:     "6380",
		Password: "",
		DB:       0,
	}

	fmt.Printf("")

	redisAdapter := redis_adapter.New(rediscfg)
	redisDB := redis_db.New(redisAdapter.Client().Conn())

	eventRepo := postgresql2.NewEventsRepository(controller)
	webhookRepo := postgresql2.NewWebhooksRepository(controller)

	eventSrv := event_service.New(redisDB, eventRepo, webhookRepo)

	return GlobalServices{
		redisDB:      redisDB,
		eventService: eventSrv,
	}
}

func migrateOperation(flag string, mg migrator.Migrator) {
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
