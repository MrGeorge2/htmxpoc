package db

import (
	"context"
	"database/sql"
	"fmt"
	"gopoc/internal/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

var dbConn *pgx.Conn

func CreateOrGetDatabase(ctx context.Context) (*pgx.Conn, error) {
	conf, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	if dbConn != nil {
		return dbConn, nil
	}

	err = createDbIfNotExists(ctx, conf)
	if err != nil {
		return nil, err
	}

	dbConn, err = createDbConnection(ctx, conf)
	if err != nil {
		return nil, err
	}

	err = migrateDb(conf)

	return dbConn, err
}

func migrateDb(conf config.Config) error {
	db, err := sql.Open("postgres", conf.Database.CreateConnectionString())

	if err != nil {
		return err
	}

	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		return err
	}

	defer driver.Close()

	m, err := migrate.NewWithDatabaseInstance(
		"file://./internal/db/migrations",
		"postgres", driver)

	if err != nil {
		return err
	}

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func createDbIfNotExists(ctx context.Context, conf config.Config) error {
	con, err := sql.Open("postgres", conf.Database.CreateConnectionString())

	if err != nil {
		return err
	}

	defer con.Close()

	err = con.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Db opened")

	var count int
	err = con.QueryRow("SELECT COUNT(1) FROM pg_database WHERE datname = $1", conf.Database.DbName).Scan(&count)
	if err != nil {
		return err
	}

	exists := count > 0

	if !exists {
		fmt.Printf("Could not find database: %s\n", conf.Database.DbName)

		_, err = con.Exec(fmt.Sprintf("CREATE DATABASE %s", conf.Database.DbName))

		if err != nil {
			return err
		}

		fmt.Printf("Database: %s created \n", conf.Database.DbName)
	}

	return nil
}

func createDbConnection(ctx context.Context, conf config.Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, conf.Database.CreateConnectionString())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
