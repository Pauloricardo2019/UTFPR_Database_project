package main

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"utfpr_db/cmd/migration/migrations"
	"utfpr_db/config"
)

func main() {
	cfg := config.GetConfig()

	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{FullSaveAssociations: true})
	if err != nil {
		log.Fatal(err)
	}

	migrationsToExec := migrations.GetMigrationsToExec()
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrationsToExec)

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")

}
