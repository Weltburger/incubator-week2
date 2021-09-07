package store

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	psgs "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

type Database struct {
	GormDB *gorm.DB
	sync.Once
}

func GetDB() *Database {
	gormDB := startDB()

	return &Database{GormDB: gormDB}
}

func startDB() *gorm.DB {
	//sqlDB, err := sql.Open("postgres", "postgres://postgres:1234@localhost:5432/simple_api?sslmode=disable")
	dbUrl := os.Getenv("DB_URL")
	sqlDB, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}


	return gormDB
}

func (s *Database) Migrate() {
	s.Do(func() {
		db, _ := s.GormDB.DB()
		driver, err := psgs.WithInstance(db, &psgs.Config{})
		fsrc, err := (&file.File{}).Open("file://migrations")
		if err != nil {
			log.Fatal(err)
		}

		m, err := migrate.NewWithInstance(
			"file",
			fsrc,
			"postgres",
			driver,
		)
		if err != nil {
			log.Fatal(err)
		}

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	})
}
